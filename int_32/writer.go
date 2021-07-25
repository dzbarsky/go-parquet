package int_32

import (
	"encoding/binary"
)

type Writer struct {
	buf []byte
	index int
}

func NewWriter(n int) *Writer {
	return &Writer{buf: make([]byte, 4*n)}
}

func (w *Writer) Write(v int32) {
	binary.LittleEndian.PutUint32(w.buf[w.index:], uint32(v))
	w.index += 4
}

func (w *Writer) Bytes() []byte {
	return w.buf
}
