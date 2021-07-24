package main

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
	"unsafe"

	"github.com/apache/thrift/lib/go/thrift"

	"parquet/float"
	"parquet/double"
	"parquet/parquet"
)

type byteCounter struct {
	n int
}

func (b *byteCounter) Write(data []byte) (int, error) {
	b.n += len(data)
	return len(data), nil
}

func write(ctx context.Context, w io.Writer, structs interface{}) error {
	rStructs := reflect.ValueOf(structs)
	nStructs := rStructs.Len()

	fileCounter := &byteCounter{}
	w = io.MultiWriter(w, fileCounter)
	_, err := w.Write([]byte(magic))
	if err != nil {
		return err
	}

	required := parquet.FieldRepetitionType_REQUIRED

	var numColumns int32
	schema := []*parquet.SchemaElement{{
		RepetitionType: &required,
		Name:           "schema",
		NumChildren:    &numColumns,
	}}

	var columns []*parquet.ColumnChunk

	// Per row-group
	totalByteSize := 0

	firstElem := rStructs.Index(0)
	ty := firstElem.Type()
	structSize := ty.Size()
	for i := 0; i < ty.NumField(); i++ {
		f := ty.Field(i)
		tag := f.Tag.Get("parquet")
		if tag == "" {
			continue
		}
		numColumns++
		firstAddr := firstElem.Field(i).UnsafeAddr()

		fieldPointer := func(idx int) unsafe.Pointer {
			return unsafe.Pointer(firstAddr + uintptr(idx)*structSize)
		}

		var data [] byte

		kind := f.Type.Kind()
		switch kind {
		case reflect.Float32:
			fw := float.NewWriter(nStructs)
			for j := 0; j < nStructs; j++ {
				fw.Write(*(*float32)(fieldPointer(j)))
			}
			data = fw.Bytes()
		case reflect.Float64:
			dw := double.NewWriter(nStructs)
			for j := 0; j < nStructs; j++ {
				dw.Write(*(*float64)(fieldPointer(j)))
			}
			data = dw.Bytes()
		default:
			return errors.New("Unhandled kind " + f.Type.Kind().String())
		}

		pageHeader := &parquet.PageHeader{
			Type:                 parquet.PageType_DATA_PAGE,
			UncompressedPageSize: int32(len(data)),
			CompressedPageSize:   int32(len(data)),
			DataPageHeader: &parquet.DataPageHeader{
				NumValues: int32(nStructs),
				Encoding:  parquet.Encoding_PLAIN,
				// TODO: fix these
				DefinitionLevelEncoding: parquet.Encoding_PLAIN,
				RepetitionLevelEncoding: parquet.Encoding_PLAIN,
			},
		}

		dataPageOffset := fileCounter.n
		err = writePageHeader(ctx, w, pageHeader)
		if err != nil {
			return err
		}

		_, err = w.Write(data)
		if err != nil {
			return err
		}

		parquetTy := parquetType(kind)
		columns = append(columns, &parquet.ColumnChunk{
			// TODO: hax, figure this out?
			FileOffset: 2,
			MetaData: &parquet.ColumnMetaData{
				Type: parquetTy,
				// TODO: hax, figure this out?
				Encodings:             []parquet.Encoding{parquet.Encoding_PLAIN},
				PathInSchema:          []string{tag},
				Codec:                 parquet.CompressionCodec_UNCOMPRESSED,
				NumValues:             int64(nStructs),
				TotalUncompressedSize: int64(len(data)),
				TotalCompressedSize:   int64(len(data)),
				DataPageOffset:        int64(dataPageOffset),
			},
		})
		schema = append(schema, &parquet.SchemaElement{
			Name:           tag,
			RepetitionType: &required,
			Type:           &parquetTy,
		})
		totalByteSize += len(data)
	}

	// TODO(zbarsky): more row groups?
	fileMetaData := &parquet.FileMetaData{
		NumRows: int64(nStructs),
		Schema:  schema,
		RowGroups: []*parquet.RowGroup{{
			Columns:       columns,
			NumRows:       int64(nStructs),
			TotalByteSize: int64(totalByteSize),
		}},
	}

	metadataLen := &byteCounter{}
	err = writeFileMetaData(ctx, io.MultiWriter(w, metadataLen), fileMetaData)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, int32(metadataLen.n))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(magic))
	return err
}

func parquetType(kind reflect.Kind) parquet.Type {
	switch kind {
	case reflect.Float32:
		return parquet.Type_FLOAT
	case reflect.Float64:
		return parquet.Type_DOUBLE
	default:
		panic("Unhandled kind " + kind.String())
	}
}

func writePageHeader(ctx context.Context, w io.Writer, header *parquet.PageHeader) error {
	proto := thrift.NewTCompactProtocol(&thrift.StreamTransport{Writer: w})
	return header.Write(ctx, proto)
}

func writeFileMetaData(ctx context.Context, w io.Writer, metadata *parquet.FileMetaData) error {
	proto := thrift.NewTCompactProtocol(&thrift.StreamTransport{Writer: w})
	return metadata.Write(ctx, proto)
}
