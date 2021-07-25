package main

import (
	"context"
	"encoding/binary"
	"errors"
	"io"
	"reflect"
	"unsafe"

	"github.com/apache/thrift/lib/go/thrift"

	"parquet/double"
	"parquet/float"
	"parquet/int_32"
	"parquet/int_64"
	"parquet/parquet"
)

type byteCounter struct {
	n int
}

func (b *byteCounter) Write(data []byte) (int, error) {
	b.n += len(data)
	return len(data), nil
}

type writeState struct {
	encodingHint map[string]parquet.Encoding
}

type writeOption func(s *writeState)

func WithEncodingHint(col string, encoding parquet.Encoding) writeOption {
	return func(s *writeState) {
		s.encodingHint[col] = encoding
	}
}

func write(ctx context.Context, w io.Writer, structs interface{}, opts ...writeOption) error {
	state := &writeState{
		encodingHint: make(map[string]parquet.Encoding),
	}
	for _, opt := range opts {
		opt(state)
	}

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

		encoding, ok := state.encodingHint[tag]
		if !ok {
			// TODO: smarter logic for deciding best encoding
			encoding = parquet.Encoding_PLAIN
		}

		kind := f.Type.Kind()
		var dataPageOffset int64
		var dictionaryPageOffset *int64

		switch encoding {
		case parquet.Encoding_PLAIN:
			var data []byte

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
			case reflect.Int32:
				iw := int_32.NewWriter(nStructs)
				for j := 0; j < nStructs; j++ {
					iw.Write(*(*int32)(fieldPointer(j)))
				}
				data = iw.Bytes()
			case reflect.Int64:
				iw := int_64.NewWriter(nStructs)
				for j := 0; j < nStructs; j++ {
					iw.Write(*(*int64)(fieldPointer(j)))
				}
				data = iw.Bytes()
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

			dataPageOffset = int64(fileCounter.n)
			err = writePageHeader(ctx, w, pageHeader)
			if err != nil {
				return err
			}

			_, err = w.Write(data)
			if err != nil {
				return err
			}
		case parquet.Encoding_PLAIN_DICTIONARY:
			var dictBytes, dataBytes []byte
			var bitWidth int8
			var nDictValues int

			switch kind {
			case reflect.Float32:
				fd := NewFloatDict(nStructs)
				for j := 0; j < nStructs; j++ {
					fd.Write(*(*float32)(fieldPointer(j)))
				}
				nDictValues = fd.NDictValues()
				dictBytes = fd.DictBytes()
				bitWidth, dataBytes = fd.DataBytes()
			default:
				return errors.New("Unhandled kind " + f.Type.Kind().String())
			}

			dictPageHeader := &parquet.PageHeader{
				Type:                 parquet.PageType_DICTIONARY_PAGE,
				UncompressedPageSize: int32(len(dictBytes)),
				CompressedPageSize:   int32(len(dictBytes)),
				DictionaryPageHeader: &parquet.DictionaryPageHeader{
					NumValues: int32(nDictValues),
					Encoding:  parquet.Encoding_PLAIN,
				},
			}

			currOffset := int64(fileCounter.n)
			dictionaryPageOffset = &currOffset
			err = writePageHeader(ctx, w, dictPageHeader)
			if err != nil {
				return err
			}

			_, err = w.Write(dictBytes)
			if err != nil {
				return err
			}

			// First byte is bitwidth, then next 4 bytes are the encoded data length
			dataPageSize := int32(len(dataBytes) + 5)
			dataPageHeader := &parquet.PageHeader{
				Type:                 parquet.PageType_DATA_PAGE,
				UncompressedPageSize: dataPageSize,
				CompressedPageSize:   dataPageSize,
				DataPageHeader: &parquet.DataPageHeader{
					NumValues: int32(nStructs),
					Encoding:  parquet.Encoding_PLAIN_DICTIONARY,
					// TODO: fix these
					DefinitionLevelEncoding: parquet.Encoding_PLAIN,
					RepetitionLevelEncoding: parquet.Encoding_PLAIN,
				},
			}

			dataPageOffset = int64(fileCounter.n)
			err = writePageHeader(ctx, w, dataPageHeader)
			if err != nil {
				return err
			}

			_, err = w.Write([]byte{byte(bitWidth)})
			if err != nil {
				return err
			}

			/*err = binary.Write(w, binary.LittleEndian, int32(len(dataBytes)))
			if err != nil {
				return err
			}*/

			_, err = w.Write(dataBytes)
			if err != nil {
				return err
			}
		default:
			panic("Unhandled encoding " + encoding.String())
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
				DataPageOffset:        dataPageOffset,
				DictionaryPageOffset:  dictionaryPageOffset,
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
	case reflect.Int32:
		return parquet.Type_INT32
	case reflect.Int64:
		return parquet.Type_INT64
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
