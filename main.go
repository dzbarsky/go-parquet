package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"unsafe"

	"github.com/DataDog/zstd"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/golang/snappy"

	"parquet/bytearray"
	"parquet/double"
	"parquet/float"
	"parquet/int_32"
	"parquet/int_64"
	"parquet/parquet"
)

type Test struct {
	Test      float64 `parquet:"test"`
	TestFloat float64 `parquet:"test_float"`
	TestInt   int64   `parquet:"test_int"`
	TestBytes []byte  `parquet:"test_str"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	must(err)

	data, err := io.ReadAll(f)
	must(err)

	parquetFile := newFile(data)
	destStructs := make([]Test, parquetFile.NumRows())
	parse(parquetFile, destStructs)
	fmt.Println(destStructs[:5])
}

func parse(f *File, destStructs interface{}) {
	ctx := context.TODO()
	r := f.r

	reps := map[string]parquet.FieldRepetitionType{}
	for _, s := range f.metadata.Schema {
		reps[s.Name] = *s.RepetitionType
	}

	previousRowGroupsTotalRows := 0
	for _, rowGroup := range f.metadata.RowGroups {
		for _, col := range rowGroup.Columns {
			fieldIndex := -1
			firstElem := reflect.ValueOf(destStructs).Index(0)
			ty := firstElem.Type()
			structSize := ty.Size()
			for i := 0; i < ty.NumField(); i++ {
				tag := ty.Field(i).Tag.Get("parquet")
				if tag == col.MetaData.PathInSchema[0] {
					fieldIndex = i
					break
				}
			}
			if fieldIndex == -1 {
				fmt.Printf("Ignoring column %s\n", col.MetaData.PathInSchema[0])
				continue
				// panic("field not found")
			}
			offset := firstElem.Field(fieldIndex).UnsafeAddr() - firstElem.UnsafeAddr()

			var dictVals interface{}
			//fmt.Println(col.MetaData)
			dictOffset := col.MetaData.DictionaryPageOffset
			if dictOffset != nil {
				_, err := r.Seek(*dictOffset, io.SeekStart)
				must(err)
				dictPageHeader, err := readPageHeader(ctx, r)
				must(err)

				dictVals = readDictPage(col.MetaData, dictPageHeader, r)
				//fmt.Println(dictVals)
			}

			dataOffset := col.MetaData.DataPageOffset
			_, err := r.Seek(dataOffset, io.SeekStart)
			must(err)
			dataPageHeader, err := readPageHeader(ctx, r)
			must(err)

			//fmt.Println(col)

			baseAddr := uintptr(firstElem.UnsafeAddr()) + uintptr(offset)
			fieldPointer := func(idx int) unsafe.Pointer {
				rowIdx := uintptr(previousRowGroupsTotalRows + idx)
				return unsafe.Pointer(baseAddr + rowIdx*structSize)
			}

			switch dataPageHeader.DataPageHeader.Encoding {
			case parquet.Encoding_PLAIN_DICTIONARY:
				defs, vals := readDataPageRLE(col.MetaData, reps[col.MetaData.PathInSchema[0]], dataPageHeader, &byteReader{r})
				//fmt.Println(col.MetaData)
				//fmt.Println(dataPageHeader)

				//fmt.Println(vals)
				//fmt.Println(dictVals)

				// reflect way is slower but safer
				// may want a hybrid approach if we get to decoding nested structures.
				//s := reflect.ValueOf(&destStructs[idx]).Elem()
				//floatVal := dictVals[v].(float64)
				//s.Field(fieldIndex).SetFloat(floatVal)

				// TODO: reading nullable values into a non-nullable field results in a 0-value
				// We make an exception for floats because pandas encodes NaN as null.
				switch col.MetaData.Type {
				case parquet.Type_FLOAT:
					values := dictVals.([]float32)
					for i, v := range vals {
						if defs[i] == 0 {
							*(*float32)(fieldPointer(i)) = float32(math.NaN())
						} else {
							*(*float32)(fieldPointer(i)) = values[v]
						}
					}
				case parquet.Type_DOUBLE:
					values := dictVals.([]float64)
					for i, v := range vals {
						if defs[i] == 0 {
							*(*float64)(fieldPointer(i)) = math.NaN()
						} else {
							*(*float64)(fieldPointer(i)) = values[v]
						}
					}
				case parquet.Type_BYTE_ARRAY:
					values := dictVals.([][]byte)
					for i, v := range vals {
						*(*[]byte)(fieldPointer(i)) = values[v]
					}
				case parquet.Type_INT32:
					values := dictVals.([]int32)
					for i, v := range vals {
						*(*int32)(fieldPointer(i)) = values[v]
					}
				case parquet.Type_INT64:
					values := dictVals.([]int64)
					for i, v := range vals {
						*(*int64)(fieldPointer(i)) = values[v]
					}
				default:
					panic("Cannot read type: " + col.MetaData.Type.String())
				}

			case parquet.Encoding_PLAIN:
				data := readDecompressed(col.MetaData, dataPageHeader, r)
				n := int(col.MetaData.NumValues);
				switch col.MetaData.Type {
				case parquet.Type_FLOAT:
					fr := float.NewReader(data)
					for i := 0; i < n; i++ {
						*(*float32)(fieldPointer(i)) = fr.Next()
					}
					must(fr.Error())
				case parquet.Type_DOUBLE:
					dr := double.NewReader(data)
					for i := 0; i < n; i++ {
						*(*float64)(fieldPointer(i)) = dr.Next()
					}
					must(dr.Error())
				case parquet.Type_INT32:
					ir := int_32.NewReader(data)
					for i := 0; i < n; i++ {
						*(*int32)(fieldPointer(i)) = ir.Next()
					}
					must(ir.Error())
				default:
					panic("Cannot read type: " + col.MetaData.Type.String())
				}
			default:
				panic("wrong encoding: " + dataPageHeader.DataPageHeader.Encoding.String())
			}
		}
		previousRowGroupsTotalRows += int(rowGroup.NumRows)
	}
}

func readPageHeader(ctx context.Context, r io.Reader) (*parquet.PageHeader, error) {
	proto := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
	pageHeader := parquet.NewPageHeader()
	err := pageHeader.Read(ctx, proto)
	return pageHeader, err
}

func wrapReader(col *parquet.ColumnMetaData, header *parquet.PageHeader, r io.Reader) io.Reader {
	switch col.Codec {
	case parquet.CompressionCodec_UNCOMPRESSED:
		return r
	default:
		decompressed := readDecompressed(col, header, r)
		return bytes.NewReader(decompressed)
	}
}

func readDecompressed(col *parquet.ColumnMetaData, header *parquet.PageHeader, r io.Reader) []byte {
	data := make([]byte, header.CompressedPageSize)
	_, err := io.ReadFull(r, data)
	must(err)

	if col.Codec == parquet.CompressionCodec_UNCOMPRESSED {
		return data
	}

	decompressed := make([]byte, header.UncompressedPageSize)

	switch col.Codec {
	case parquet.CompressionCodec_ZSTD:
		decompressed, err = zstd.Decompress(decompressed, data)
	case parquet.CompressionCodec_SNAPPY:
		decompressed, err = snappy.Decode(decompressed, data)
	default:
		panic("Unsupported compression: " + col.Codec.String())
	}
	must(err)
	return decompressed
}

func readDictPage(col *parquet.ColumnMetaData, header *parquet.PageHeader, r io.Reader) interface{} {
	if header.Type != parquet.PageType_DICTIONARY_PAGE {
		panic("wrong page type")
	}

	data := readDecompressed(col, header, r)

	num := header.DictionaryPageHeader.NumValues

	switch col.Type {
	case parquet.Type_FLOAT:
		vals := make([]float32, num)
		fr := float.NewReader(data)
		for i := range vals {
			vals[i] = fr.Next()
		}
		must(fr.Error())
		return vals
	case parquet.Type_DOUBLE:
		vals := make([]float64, num)
		dr := double.NewReader(data)
		for i := range vals {
			vals[i] = dr.Next()
		}
		must(dr.Error())
		return vals
	case parquet.Type_INT32:
		vals := make([]int32, num)
		ir := int_32.NewReader(data)
		for i := range vals {
			vals[i] = ir.Next()
		}
		must(ir.Error())
		return vals
	case parquet.Type_INT64:
		vals := make([]int64, num)
		ir := int_64.NewReader(data)
		for i := range vals {
			vals[i] = ir.Next()
		}
		must(ir.Error())
		return vals
	case parquet.Type_BYTE_ARRAY:
		vals := make([][]byte, num)
		bar := bytearray.NewReader(data)
		for i := range vals {
			vals[i] = bar.Next()
		}
		must(bar.Error())
		return vals
	default:
		panic("Cannot read type: " + col.Type.String())
	}
}

// readDataPageRLE returns definition levels and values
func readDataPageRLE(
	col *parquet.ColumnMetaData,
	repType parquet.FieldRepetitionType,
	header *parquet.PageHeader,
	r reader,
) ([]int32, []int32) {
	if header.Type != parquet.PageType_DATA_PAGE {
		panic("wrong page type")
	}

	r = &byteReader{wrapReader(col, header, r)}

	switch header.DataPageHeader.Encoding {
	case parquet.Encoding_PLAIN_DICTIONARY, parquet.Encoding_PLAIN:
	default:
		panic("wrong encoding: " + header.DataPageHeader.Encoding.String())
	}

	buf2 := make([]byte, header.UncompressedPageSize)
	_, err2 := io.ReadFull(r, buf2)
	must(err2)
	//fmt.Println(buf2)
	r = bytes.NewReader(buf2)

	switch repType {
	case parquet.FieldRepetitionType_REQUIRED:
		buf := make([]byte, 1)
		_, err := io.ReadFull(r, buf)
		must(err)
		bitWidth := int(buf[0])

		vals := make([]int32, header.DataPageHeader.NumValues)
		hr := &hybridReader{r: r, bitWidth: bitWidth}
		for i := range vals {
			val, err := hr.Next()
			must(err)
			vals[i] = val
		}
		return nil, vals
	case parquet.FieldRepetitionType_OPTIONAL:
		defs := make([]int32, 0, header.DataPageHeader.NumValues)
		//fmt.Println(header)
		var size uint32
		err := binary.Read(r, binary.LittleEndian, &size)
		must(err)
		//fmt.Println("size: ", size)

		// TODO: This might be repition level actually
		// TODO: Figure out where to compute max def level
		// TODO: need these to handle nullability correctly?
		hr := &hybridReader{
			r:        &byteReader{io.LimitReader(r, int64(size))},
			bitWidth: 1, // max definition level = 1
		}

		for {
			v, err := hr.Next()
			if err == io.EOF {
				break
			}
			defs = append(defs, v)
			if err != nil {
				panic(err)
			}
		}

		buf := make([]byte, 1)
		_, err = io.ReadFull(r, buf)
		must(err)
		bitWidth := int(buf[0])

		vals := make([]int32, header.DataPageHeader.NumValues)
		hr = &hybridReader{r: r, bitWidth: bitWidth}
		for i, defined := range defs {
			if defined == 0 {
				continue
			}
			val, err := hr.Next()
			must(err)
			vals[i] = val
		}

		return defs, vals
	default:
		panic("Unsupported")
	}
}

type byteReader struct {
	io.Reader
}

func (b *byteReader) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	// TODO: io.ReadFull would be safer
	_, err := b.Reader.Read(buf)
	return buf[0], err
}
