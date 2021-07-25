package int_64

import (
	"encoding/binary"
)

type Writer struct {
	buf   []byte
	index int
}

func NewWriter(n int) *Writer {
	return &Writer{buf: make([]byte, 8*n)}
}

func (w *Writer) Write(v int64) {
	binary.LittleEndian.PutUint64(w.buf[w.index:], uint64(v))
	w.index += 8
}

func (w *Writer) Bytes() []byte {
	return w.buf
}
