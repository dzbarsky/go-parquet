package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type hybridWriter struct {
	w        io.Writer
	bitWidth int

	// TODO: we aren't doing any bitpacking yet

	currRLEVal    int32
	currRLELength int
	scratch       [4]byte
}

func newHybridWriter(w io.Writer, bitWidth int) *hybridWriter {
	fmt.Println("write bw", bitWidth)
	return &hybridWriter{
		w:        w,
		bitWidth: bitWidth,
	}
}

func (hw *hybridWriter) Write(v int32) error {
	if v == hw.currRLEVal {
		hw.currRLELength += 1
		return nil
	}

	n := binary.PutUvarint(hw.scratch[:], uint64(hw.currRLELength<<1))
	_, err := hw.w.Write(hw.scratch[:n])
	if err != nil {
		return err
	}

	buf := make([]byte, (hw.bitWidth+7)/8)
	switch len(buf) {
	case 1:
		buf[0] = byte(hw.currRLEVal)
	default:
		panic("Bad int size")
	}

	_, err = hw.w.Write(buf)
	if err != nil {
		return err
	}

	hw.currRLEVal = v
	return nil
}
