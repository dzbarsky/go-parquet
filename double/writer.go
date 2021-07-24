package double

import (
	"encoding/binary"
	"io"
	"math"
)

type Writer struct {
	w io.Writer
	buf [8]byte
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(v float64) error {
	binary.LittleEndian.PutUint64(w.buf[:], math.Float64bits(v))
	_, err := w.w.Write(w.buf[:])
	return err
}
