package double

import (
	"encoding/binary"
	"math"
)

type Writer struct {
	buf []byte
	index int
}

func NewWriter(n int) *Writer {
	return &Writer{buf: make([]byte, 8*n)}
}

func (w *Writer) Write(v float64) {
	binary.LittleEndian.PutUint64(w.buf[w.index:], math.Float64bits(v))
	w.index += 8
}

func (w *Writer) Bytes() []byte {
	return w.buf
}
