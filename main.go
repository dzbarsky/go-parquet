package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"reflect"
	"unsafe"

	"github.com/apache/thrift/lib/go/thrift"

	"parquet/double"
	"parquet/float"
	"parquet/parquet"
)

type Test struct {
	Test float64 `parquet:"test"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func isMagic(d []byte) bool {
	return string(d) == "PAR1"
}

func main() {
	f, err := os.Open(os.Args[1])
	must(err)

	data, err := io.ReadAll(f)
	must(err)

	fmt.Println(parse(data))
}

func parse(data []byte) []Test {
	ctx := context.TODO()

	if !isMagic(data[:4]) || !isMagic(data[len(data)-4:]) {
		panic("Not a parquet file")
	}

	// Read footer size from end
	r := bytes.NewReader(data)
	_, err := r.Seek(-8, io.SeekEnd)
	must(err)
	sizeBuf := make([]byte, 4)
	_, err = r.Read(sizeBuf)
	must(err)
	size := binary.LittleEndian.Uint32(sizeBuf)

	// Read footer
	_, err = r.Seek(-8-int64(size), io.SeekEnd)
	must(err)
	footerReader := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
	fileMD := parquet.NewFileMetaData()
	err = fileMD.Read(ctx, footerReader)
	must(err)
	destStructs := make([]Test, fileMD.NumRows)
	previousRowGroupsTotalRows := 0
	for _, rowGroup := range fileMD.RowGroups {
		for _, col := range rowGroup.Columns {
			fieldIndex := -1
			ty := reflect.TypeOf(destStructs[0])
			for i := 0; i < ty.NumField(); i++ {
				tag := ty.Field(i).Tag.Get("parquet")
				if tag == col.MetaData.PathInSchema[0] {
					fieldIndex = i
					break
				}
			}
			if fieldIndex == -1 {
				panic("field not found")
			}
			s := reflect.ValueOf(&destStructs[0]).Elem()
			addr := s.UnsafeAddr()
			addr2 := s.Field(fieldIndex).UnsafeAddr()
			offset := addr2 - addr

			var dictVals []interface{}
			//fmt.Println(col.MetaData)
			dictOffset := col.MetaData.DictionaryPageOffset
			if dictOffset != nil {
				//fmt.Println("Will read dict from", *dictOffset)
				_, err := r.Seek(*dictOffset, io.SeekStart)
				must(err)
				dictPageHeader, err := readPageHeader(ctx, r)
				must(err)

				dictVals = readDictPage(col.MetaData, dictPageHeader, r)
				//fmt.Println(dictVals)
			}

			dataOffset := col.MetaData.DataPageOffset
			//fmt.Println("Will read data from", dataOffset)
			_, err := r.Seek(dataOffset, io.SeekStart)
			must(err)
			dataPageHeader, err := readPageHeader(ctx, r)
			must(err)

			vals := readDataPage(dataPageHeader, r, dictVals)

			for i, v := range vals {
				idx := previousRowGroupsTotalRows + i
				floatVal := dictVals[v.(int32)].(float64)
				// reflect way is slower but safer
				// may want a hybrid approach if we get to decoding nested structures.
				//s := reflect.ValueOf(&destStructs[idx]).Elem()
				//s.Field(fieldIndex).SetFloat(floatVal)

				newP := unsafe.Pointer(uintptr(unsafe.Pointer(&destStructs[idx])) + uintptr(offset))
				*(*float64)(newP) = floatVal
			}
		}
		previousRowGroupsTotalRows += int(rowGroup.NumRows)
	}
	return destStructs
}

func readPageHeader(ctx context.Context, r io.Reader) (*parquet.PageHeader, error) {
	proto := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
	pageHeader := parquet.NewPageHeader()
	err := pageHeader.Read(ctx, proto)
	return pageHeader, err
}

func readDictPage(col *parquet.ColumnMetaData, header *parquet.PageHeader, r io.Reader) []interface{} {
	if header.Type != parquet.PageType_DICTIONARY_PAGE {
		panic("wrong page type")
	}

	vals := make([]interface{}, header.DictionaryPageHeader.NumValues)

	switch col.Type {
	case parquet.Type_FLOAT:
		fr := float.NewReader(r)
		for i := range vals {
			vals[i] = fr.Next()
		}
		must(fr.Error())
	case parquet.Type_DOUBLE:
		dr := double.NewReader(r)
		for i := range vals {
			vals[i] = dr.Next()
		}
		must(dr.Error())
	default:
		panic("Cannot read type: " + col.Type.String())
	}
	return vals
}

func readDataPage(header *parquet.PageHeader, r reader, dictVals []interface{}) []interface{} {
	//fmt.Println(header)
	if header.Type != parquet.PageType_DATA_PAGE {
		panic("wrong page type")
	}
	if header.DataPageHeader.Encoding != parquet.Encoding_PLAIN_DICTIONARY {
		panic("wrong encoding")
	}
	vals := make([]interface{}, 0, header.DataPageHeader.NumValues)
	//fmt.Println(header)

	buf2 := make([]byte, header.UncompressedPageSize)
	_, err2 := io.ReadFull(r, buf2)
	must(err2)
	//fmt.Println(buf2)
	r = bytes.NewReader(buf2)

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
		_, err := hr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}

	buf := make([]byte, 1)
	_, err = r.Read(buf)
	must(err)
	bitWidth := int(buf[0])

	hr = &hybridReader{r: r, bitWidth: bitWidth}
	for {
		val, err := hr.Next()
		if err == io.EOF {
			break
		}
		must(err)
		vals = append(vals, val)
	}

	return vals
}

type byteReader struct {
	io.Reader
}

func (b *byteReader) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	_, err := b.Reader.Read(buf)
	return buf[0], err
}
