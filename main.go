package main

import (
	"os"
	"encoding/binary"
	"bytes"
	"fmt"
	"context"
	"io"

	"github.com/apache/thrift/lib/go/thrift"

	"parquet/double"
	"parquet/float"
	"parquet/parquet"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func isMagic(d []byte) bool {
	return string(d) == "PAR1"
}

func main() {
	ctx := context.TODO()

	f, err := os.Open(os.Args[1])
	must(err)

	data, err := io.ReadAll(f)
	must(err)

	if !isMagic(data[:4]) || !isMagic(data[len(data)-4:]) {
		panic("Not a parquet file")
	}

	// Read footer size from end
	r := bytes.NewReader(data)
	_, err = r.Seek(-8, io.SeekEnd)
	must(err)
	sizeBuf := make([]byte, 4)
	_, err = r.Read(sizeBuf)
	must(err)
	size := binary.LittleEndian.Uint32(sizeBuf)

	// Read footer
	_, err = r.Seek(-8 - int64(size), io.SeekEnd)
	must(err)
	footerReader := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
	fileMD := parquet.NewFileMetaData()
       	err = fileMD.Read(ctx, footerReader)
	must(err)
	for _, s := range fileMD.Schema {
		_ = s
		//fmt.Println(s)
	}
	for _, rowGroup := range fileMD.RowGroups {
		for _, col := range rowGroup.Columns {
			var dictVals []interface{}
			fmt.Println(col.MetaData)
			dictOffset := col.MetaData.DictionaryPageOffset
			if dictOffset != nil {
				fmt.Println("Will read dict from", *dictOffset)
				_, err := r.Seek(*dictOffset, io.SeekStart)
				must(err)
				dictPageHeader, err := readPageHeader(ctx, r)
				must(err)

				dictVals = readDictPage(col.MetaData, dictPageHeader, r)
				//fmt.Println(dictVals)
			}

			dataOffset := col.MetaData.DataPageOffset
			fmt.Println("Will read data from", dataOffset)
			_, err := r.Seek(dataOffset, io.SeekStart)
			must(err)
			dataPageHeader, err := readPageHeader(ctx, r)
			must(err)

			vals := readDataPage(dataPageHeader, r,  dictVals)
	
			out := make([]float64, len(vals))
			for i, v := range vals {
				out[i] = dictVals[v.(int32)].(float64)
			}
			fmt.Println(out)
		}
	}
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
	fmt.Println(header)
	if header.Type != parquet.PageType_DATA_PAGE {
		panic("wrong page type")
	}
	if header.DataPageHeader.Encoding != parquet.Encoding_PLAIN_DICTIONARY {
		panic("wrong encoding")
	}
	vals := make([]interface{}, header.DataPageHeader.NumValues)
	fmt.Println(header)

	buf2 := make([]byte, header.UncompressedPageSize)
	_, err2 := io.ReadFull(r, buf2)
	must(err2)
	fmt.Println(buf2)
	r = bytes.NewReader(buf2)

	var size uint32
	err := binary.Read(r, binary.LittleEndian, &size)
	must(err)
	fmt.Println("size: ", size)

	// TODO: This might be repition level actually
	// TODO: Figure out where to compute max def level
	// TODO: need these to handle nullability correctly?
	readHybrid(&byteReader{io.LimitReader(r, int64(size))},
		1, nil) // max definition level = 1

	fmt.Println("Real r pre", r)
	buf := make([]byte, 1)
	_, err = r.Read(buf)
	must(err)
	bitWidth := int(buf[0])
	fmt.Println("Will read with bitwidth ", bitWidth)
	fmt.Println("Real r post", r)
	readHybrid(r, bitWidth, vals[:0])
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

func (b* byteReader) print() {
	fmt.Println(b.Reader)
}

type reader interface {
	io.Reader
	io.ByteReader
}

func readHybrid(r reader, bitWidth int, vals []interface{}) {
	for {
		header, err := binary.ReadUvarint(r)
		if err == io.EOF {
			return
		}
		fmt.Println("head", header)
		if (header & 1) != 0{
			bitPackedRunLen := (header >> 1)

			for i := uint64(0); i < bitPackedRunLen; i++ {
				scratch := make([]byte, bitWidth)
				_, err = r.Read(scratch)
				must(err)

				var unpacked [8]int32
				switch bitWidth {
				case 3:
					unpacked = unpack8int32_3(scratch)
				case 4:
					unpacked = unpack8int32_4(scratch)
				default:
					panic("Unhandled bitwidth")
				}
				for _, b := range unpacked {
					vals = append(vals, b)
				}
			}
		} else {
			rleRunLen := header >> 1
			buf := make([]byte, (bitWidth + 7) / 8)
			_, err = r.Read(buf)
			must(err)
			if len(buf) != 1 {
				panic("need to read bigger ints")
			}
			repeatedVal := buf[0]

			fmt.Printf("run len: %v, val: %v\n", rleRunLen, repeatedVal)
			for j := uint64(0); j < rleRunLen; j++ {
				vals = append(vals, int32(repeatedVal))
			}
		}
	}
}

func unpack8int32_3(data []byte) (out [8]int32) {
	out[0] = int32(uint32((data[0]>>0)&7) << 0)
	out[1] = int32(uint32((data[0]>>3)&7) << 0)
	out[2] = int32(uint32((data[0]>>6)&3) << 0 | uint32((data[1]>>0)&1)<<2)
	out[3] = int32(uint32((data[1]>>1)&7) << 0)
	out[4] = int32(uint32((data[1]>>4)&7) << 0)
	out[5] = int32(uint32((data[1]>>7)&1) << 0 | uint32((data[2]>>0)&3)<<1)
	out[6] = int32(uint32((data[2]>>2)&7) << 0)
	out[7] = int32(uint32((data[2]>>5)&7) << 0)
	return
}

func unpack8int32_4(data []byte) (out [8]int32) {
	out[0] = int32(uint32((data[0]>>0)&15) << 0)
	out[1] = int32(uint32((data[0]>>4)&15) << 0)
	out[2] = int32(uint32((data[1]>>0)&15) << 0)
	out[3] = int32(uint32((data[1]>>4)&15) << 0)
	out[4] = int32(uint32((data[2]>>0)&15) << 0)
	out[5] = int32(uint32((data[2]>>4)&15) << 0)
	out[6] = int32(uint32((data[3]>>0)&15) << 0)
	out[7] = int32(uint32((data[3]>>4)&15) << 0)
	return
}
