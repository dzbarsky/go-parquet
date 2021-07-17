package main

import (
	"encoding/binary"
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
	unpacked     []int32
}

func (hr *hybridReader) Next() (interface{}, error) {
	if hr.rleRemaining > 0 {
		//fmt.Printf("run length: %v, val: %v\n", hr.rleRemaining, hr.rleValue)
		hr.rleRemaining -= 1
		return hr.rleValue, nil
	}

	if len(hr.unpacked) > 0 {
		//fmt.Printf("unpacked: %v\n", hr.unpacked)
		ret := hr.unpacked[0]
		hr.unpacked = hr.unpacked[1:]
		return ret, nil
	}

	err := hr.readMore()
	if err != nil {
		return nil, err
	}
	return hr.Next()
}

func (hr *hybridReader) readMore() error {
	header, err := binary.ReadUvarint(hr.r)
	if err == io.EOF {
		return err
	}
	//fmt.Println("head", header)
	if (header & 1) != 0 {
		bitPackedRunLen := (header >> 1)

		hr.unpacked = nil
		for i := uint64(0); i < bitPackedRunLen; i++ {
			scratch := make([]byte, hr.bitWidth)
			_, err = hr.r.Read(scratch)
			if err != nil {
				return err
			}

			var unpacked [8]int32
			switch hr.bitWidth {
			case 3:
				unpacked = unpack8int32_3(scratch)
			case 4:
				unpacked = unpack8int32_4(scratch)
			default:
				panic("Unhandled bitwidth")
			}
			hr.unpacked = append(hr.unpacked, unpacked[:]...)
		}
	} else {
		hr.rleRemaining = header >> 1
		buf := make([]byte, (hr.bitWidth+7)/8)
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
