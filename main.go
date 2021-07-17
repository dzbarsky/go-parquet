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
				fmt.Println(dictVals)
			}

			dataOffset := col.MetaData.DataPageOffset
			fmt.Println("Will read data from", dataOffset)
			_, err := r.Seek(dataOffset, io.SeekStart)
			must(err)
			dataPageHeader, err := readPageHeader(ctx, r)
			must(err)

			vals := readDataPage(dataPageHeader, r,  dictVals)
			fmt.Println(vals[:10])
			break
	
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

func readDataPage(header *parquet.PageHeader, r io.Reader, dictVals []interface{}) []interface{} {
	if header.Type != parquet.PageType_DATA_PAGE {
		panic("wrong page type")
	}
	if header.DataPageHeader.Encoding != parquet.Encoding_PLAIN_DICTIONARY {
		panic("wrong encoding")
	}
	vals := make([]interface{}, header.DataPageHeader.NumValues)
	fmt.Println(header)

	fmt.Println(r)
	
	var size uint32
	err := binary.Read(r, binary.LittleEndian, &size)
	must(err)
	fmt.Println("size: ", size)

	bitWidth := 3
	/*buf := make([]byte, 1)
	_, err = r.Read(buf)
	must(err)
	bitWidth = int(buf[0])
	fmt.Println("bitwdith: ", bitWidth)*/

	scratch := make([]byte, bitWidth)
	_, err = r.Read(scratch)
	must(err)
	unpacked := unpack8int32_3(scratch)
	fmt.Println(unpacked)

	return vals
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
