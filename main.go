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
			fmt.Println(col.MetaData)
			dictOffset := col.MetaData.DictionaryPageOffset
			if dictOffset != nil {
				fmt.Println("Will read dict from", *dictOffset)
				_, err := r.Seek(*dictOffset, io.SeekStart)
				must(err)
				dictPageHeader, err := readPageHeader(ctx, r, col.MetaData.GetTotalCompressedSize())
				must(err)

				fmt.Println(dictPageHeader)
				vals := readDictPage(col.MetaData, dictPageHeader, r)
				fmt.Println(vals)
			}

			dataOffset := col.MetaData.DataPageOffset
			fmt.Println("Will read data from", dataOffset)
			_, err := r.Seek(dataOffset, io.SeekStart)
			must(err)
			dataPageHeader, err := readPageHeader(ctx, r, col.MetaData.GetTotalCompressedSize())
			must(err)
			fmt.Println(dataPageHeader)

			break
	
		}
	}


	/*buffer := bytes.NewBuffer(data)
	buf := &thrift.TMemoryBuffer{
		Buffer: buffer,
	}
	thriftReader := thrift.NewTCompactProtocol(buf)

       	pageHeader := parquet.NewPageHeader()
       	err = pageHeader.Read(context.TODO(), thriftReader)
	must(err)
	fmt.Println(pageHeader)

	switch pageHeader.Type {
	case parquet.PageType_DICTIONARY_PAGE:
		out := make([]byte, pageHeader.CompressedPageSize)
		_, _ = buffer.Read(out)
	case parquet.PageType_DATA_PAGE:
		out := make([]byte, pageHeader.CompressedPageSize)
		_, _ = buffer.Read(out)
	}*/
}

func readPageHeader(ctx context.Context, r io.Reader, size int64) (*parquet.PageHeader, error) {
	proto := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
       	pageHeader := parquet.NewPageHeader()
       	err := pageHeader.Read(ctx, proto)
	return pageHeader, err
}

func readDictPage(col *parquet.ColumnMetaData, header *parquet.PageHeader, r io.Reader) []interface{} {
	fmt.Println(col)
	fmt.Println(header)

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
