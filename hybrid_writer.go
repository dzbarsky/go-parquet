package main

import (
	"encoding/binary"
)

type hybridWriter struct {
	buf []byte
	offset int

	bitWidthNBytes int

	// TODO: we aren't doing any bitpacking yet

	currRLEVal    int32
	currRLELength int
	scratch       [4]byte
}

func newHybridWriter(nValues, bitWidth int) *hybridWriter {
	bitWidthNBytes := (bitWidth+7)/8
	return &hybridWriter{
		buf: make([]byte, nValues * (binary.MaxVarintLen64 + bitWidthNBytes)),
		bitWidthNBytes: bitWidthNBytes,
	}
}

func (hw *hybridWriter) Write(v int32) {
	if v == hw.currRLEVal {
		hw.currRLELength += 1
		return
	}

	hw.flush()
	hw.currRLEVal = v
}

func (hw *hybridWriter) Flush() []byte {
	hw.flush()
	return hw.buf[:hw.offset]
}

func (hw *hybridWriter) flush() {
	n := binary.PutUvarint(hw.buf[hw.offset:], uint64(hw.currRLELength<<1))
	hw.offset += n

	switch hw.bitWidthNBytes {
	case 1:
		hw.buf[hw.offset] = byte(hw.currRLEVal)
	case 2:
		hw.buf[hw.offset]  = byte(hw.currRLEVal)
		hw.buf[hw.offset+1] = byte(hw.currRLEVal>>8)
	case 3:
		hw.buf[hw.offset]  = byte(hw.currRLEVal)
		hw.buf[hw.offset+1] = byte(hw.currRLEVal>>8)
		hw.buf[hw.offset+2] = byte(hw.currRLEVal>>16)
	case 4:
		hw.buf[hw.offset]  = byte(hw.currRLEVal)
		hw.buf[hw.offset+1] = byte(hw.currRLEVal>>8)
		hw.buf[hw.offset+2] = byte(hw.currRLEVal>>16)
		hw.buf[hw.offset+3] = byte(hw.currRLEVal>>24)
	default:
		panic("Bad int size")
	}

	hw.offset += hw.bitWidthNBytes
}
