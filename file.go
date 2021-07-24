package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"

	"github.com/apache/thrift/lib/go/thrift"

	"parquet/parquet"
)

const magic = "PAR1"

func isMagic(d []byte) bool {
	return string(d) == magic
}

type File struct {
	r        io.ReadSeeker
	metadata *parquet.FileMetaData
}

func newFile(data []byte) *File {
	if !isMagic(data[:4]) || !isMagic(data[len(data)-4:]) {
		panic("Not a parquet file")
	}

	// Read footer size from end
	r := bytes.NewReader(data)
	_, err := r.Seek(-8, io.SeekEnd)
	must(err)
	sizeBuf := make([]byte, 4)
	_, err = io.ReadFull(r, sizeBuf)
	must(err)
	size := binary.LittleEndian.Uint32(sizeBuf)

	// Read footer
	_, err = r.Seek(-8-int64(size), io.SeekEnd)
	must(err)
	footerReader := thrift.NewTCompactProtocol(&thrift.StreamTransport{Reader: r})
	fileMD := parquet.NewFileMetaData()
	err = fileMD.Read(context.TODO(), footerReader)
	must(err)

	return &File{r, fileMD}
}

func (f *File) NumRows() int64 {
	return f.metadata.NumRows
}
