package float

import (
	"encoding/binary"
	"io"
	"math"
)

type Writer struct {
	w io.Writer
	buf [4]byte
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(v float32) error {
	binary.LittleEndian.PutUint32(w.buf[:], math.Float32bits(v))
	_, err := w.w.Write(w.buf[:])
	return err
}
