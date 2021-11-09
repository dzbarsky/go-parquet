package main

import (
	"encoding/binary"
)

type hybridWriter struct {
	buf    []byte
	offset int

	bitWidthNBytes int

	// TODO: we aren't doing any bitpacking yet

	currRLEVal    int32
	currRLELength int
	scratch       [4]byte
}

func newHybridWriter(nValues, bitWidth int) *hybridWriter {
	bitWidthNBytes := (bitWidth + 7) / 8
	return &hybridWriter{
		buf:            make([]byte, nValues*(binary.MaxVarintLen64+bitWidthNBytes)),
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

	buf := hw.buf[hw.offset:]

	switch hw.bitWidthNBytes {
	case 1:
		buf[0] = byte(hw.currRLEVal)
	case 2:
		_ = buf[1]
		buf[0] = byte(hw.currRLEVal)
		buf[1] = byte(hw.currRLEVal >> 8)
	case 3:
		_ = buf[2]
		buf[0] = byte(hw.currRLEVal)
		buf[1] = byte(hw.currRLEVal >> 8)
		buf[2] = byte(hw.currRLEVal >> 16)
	case 4:
		_ = buf[3]
		buf[0] = byte(hw.currRLEVal)
		buf[1] = byte(hw.currRLEVal >> 8)
		buf[2] = byte(hw.currRLEVal >> 16)
		buf[3] = byte(hw.currRLEVal >> 24)
	default:
		panic("Bad int size")
	}

	hw.offset += hw.bitWidthNBytes
}
