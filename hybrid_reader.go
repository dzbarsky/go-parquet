package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type reader interface {
	io.Reader
	io.ByteReader
}

type hybridReader struct {
	r        reader
	bitWidth int

	rleRemaining uint64
	rleValue     int32

	bitPackedRemaining uint64
	unpacked           []int32

	// Scratch buffers. These can only be used within a single function call of hybridReader.
	unpackedScratch [8]int32
	scratch         []byte
}

func (hr *hybridReader) Next() (int32, error) {
	if hr.rleRemaining > 0 {
		//fmt.Printf("run length: %v, val: %v\n", hr.rleRemaining, hr.rleValue)
		hr.rleRemaining -= 1
		return hr.rleValue, nil
	}

	if len(hr.unpacked) > 0 {
		return hr.nextUnpackedValue(), nil
	}

	if hr.bitPackedRemaining > 0 {
		err := hr.read8BitPackedValues()
		if err != nil {
			return 0, err
		}
		return hr.nextUnpackedValue(), nil
	}

	err := hr.readMore()
	if err != nil {
		return 0, err
	}
	return hr.Next()
}

func (hr *hybridReader) nextUnpackedValue() int32 {
	ret := hr.unpacked[0]
	hr.unpacked = hr.unpacked[1:]
	return ret
}

func (hr *hybridReader) read8BitPackedValues() error {
	if cap(hr.scratch) < hr.bitWidth {
		hr.scratch = make([]byte, hr.bitWidth)
	}

	// TODO: io.ReadFull would be safer
	_, err := hr.r.Read(hr.scratch)
	if err != nil {
		return err
	}

	switch hr.bitWidth {
	case 1:
		hr.unpackedScratch = unpack8int32_1(hr.scratch)
	case 2:
		hr.unpackedScratch = unpack8int32_2(hr.scratch)
	case 3:
		hr.unpackedScratch = unpack8int32_3(hr.scratch)
	case 4:
		hr.unpackedScratch = unpack8int32_4(hr.scratch)
	case 5:
		hr.unpackedScratch = unpack8int32_5(hr.scratch)
	case 6:
		hr.unpackedScratch = unpack8int32_6(hr.scratch)
	case 7:
		hr.unpackedScratch = unpack8int32_7(hr.scratch)
	case 8:
		hr.unpackedScratch = unpack8int32_8(hr.scratch)
	case 9:
		hr.unpackedScratch = unpack8int32_9(hr.scratch)
	case 10:
		hr.unpackedScratch = unpack8int32_10(hr.scratch)
	case 11:
		hr.unpackedScratch = unpack8int32_11(hr.scratch)
	case 12:
		hr.unpackedScratch = unpack8int32_12(hr.scratch)
	case 13:
		hr.unpackedScratch = unpack8int32_13(hr.scratch)
	case 14:
		hr.unpackedScratch = unpack8int32_14(hr.scratch)
	case 15:
		hr.unpackedScratch = unpack8int32_15(hr.scratch)
	default:
		panic(fmt.Sprintf("Unhandled bitwidth %v", hr.bitWidth))
	}
	hr.unpacked = hr.unpackedScratch[:]
	hr.bitPackedRemaining -= 1
	return nil
}

func (hr *hybridReader) readMore() error {
	header, err := binary.ReadUvarint(hr.r)
	if err == io.EOF {
		return err
	}
	//fmt.Println("head", header)
	if (header & 1) != 0 {
		hr.bitPackedRemaining = (header >> 1)
		return hr.read8BitPackedValues()
	} else {
		hr.rleRemaining = header >> 1
		buf := make([]byte, (hr.bitWidth+7)/8)
		// TODO: io.ReadFull would be safer
		_, err = hr.r.Read(buf)
		if err != nil {
			return err
		}
		switch len(buf) {
		case 1:
			hr.rleValue = int32(buf[0])
		case 2:
			hr.rleValue = int32(buf[0]) + int32(buf[1])<<8
		case 3:
			hr.rleValue = int32(buf[0]) + int32(buf[1])<<8 + int32(buf[2])<<16
		case 4:
			hr.rleValue = int32(buf[0]) + int32(buf[1])<<8 + int32(buf[2])<<16 + int32(buf[3])<<24
		default:
			panic("Bad int size")
		}
	}
	return nil
}
