package float

import (
	"encoding/binary"
	"io"
	"math"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Write(v float32) error {
	return binary.Write(w.w, binary.LittleEndian, math.Float32bits(v))
}
