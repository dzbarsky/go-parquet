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
	case 3:
		hr.unpackedScratch = unpack8int32_3(hr.scratch)
	case 4:
		hr.unpackedScratch = unpack8int32_4(hr.scratch)
	case 14:
		hr.unpackedScratch = unpack8int32_14(hr.scratch)
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
		if len(buf) != 1 {
			panic("need to read bigger ints")
		}
		hr.rleValue = int32(buf[0])
	}
	return nil
}

func unpack8int32_1(data []byte) (out [8]int32) {
	_ = data[0]
	out[0] = int32(uint32((data[0]>>0)&1) << 0)
	out[1] = int32(uint32((data[0]>>1)&1) << 0)
	out[2] = int32(uint32((data[0]>>2)&1) << 0)
	out[3] = int32(uint32((data[0]>>3)&1) << 0)
	out[4] = int32(uint32((data[0]>>4)&1) << 0)
	out[5] = int32(uint32((data[0]>>5)&1) << 0)
	out[6] = int32(uint32((data[0]>>6)&1) << 0)
	out[7] = int32(uint32((data[0]>>7)&1) << 0)
	return
}

func unpack8int32_3(data []byte) (out [8]int32) {
	out[0] = int32(uint32((data[0]>>0)&7) << 0)
	out[1] = int32(uint32((data[0]>>3)&7) << 0)
	out[2] = int32(uint32((data[0]>>6)&3)<<0 | uint32((data[1]>>0)&1)<<2)
	out[3] = int32(uint32((data[1]>>1)&7) << 0)
	out[4] = int32(uint32((data[1]>>4)&7) << 0)
	out[5] = int32(uint32((data[1]>>7)&1)<<0 | uint32((data[2]>>0)&3)<<1)
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

func unpack8int32_14(data []byte) (out [8]int32) {
	out[0] = int32(uint32((data[0]>>0)&255)<<0 | uint32((data[1]>>0)&63)<<8)
	out[1] = int32(uint32((data[1]>>6)&3)<<0 | uint32((data[2]>>0)&255)<<2 | uint32((data[3]>>0)&15)<<10)
	out[2] = int32(uint32((data[3]>>4)&15)<<0 | uint32((data[4]>>0)&255)<<4 | uint32((data[5]>>0)&3)<<12)
	out[3] = int32(uint32((data[5]>>2)&63)<<0 | uint32((data[6]>>0)&255)<<6)
	out[4] = int32(uint32((data[7]>>0)&255)<<0 | uint32((data[8]>>0)&63)<<8)
	out[5] = int32(uint32((data[8]>>6)&3)<<0 | uint32((data[9]>>0)&255)<<2 | uint32((data[10]>>0)&15)<<10)
	out[6] = int32(uint32((data[10]>>4)&15)<<0 | uint32((data[11]>>0)&255)<<4 | uint32((data[12]>>0)&3)<<12)
	out[7] = int32(uint32((data[12]>>2)&63)<<0 | uint32((data[13]>>0)&255)<<6)
	return
}
