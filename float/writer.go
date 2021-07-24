package float

import (
	"encoding/binary"
	"math"
)

type Writer struct {
	buf []byte
	index int
}

func NewWriter(n int) *Writer {
	return &Writer{buf: make([]byte, 4*n)}
}

func (w *Writer) Write(v float32) {
	binary.LittleEndian.PutUint32(w.buf[w.index:], math.Float32bits(v))
	w.index += 4
}

func (w *Writer) Bytes() []byte {
	return w.buf
}
