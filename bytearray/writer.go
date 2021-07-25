package bytearray

import (
	"encoding/binary"
)

type Writer struct {
	buf    []byte
	offset int
}

func NewWriter(totalLen, nValues int) *Writer {
	return &Writer{buf: make([]byte, totalLen+4*nValues)}
}

func (w *Writer) Write(v []byte) {
	binary.LittleEndian.PutUint32(w.buf[w.offset:], uint32(len(v)))
	w.offset += 4

	copy(w.buf[w.offset:], v)
	w.offset += len(v)
}

func (w *Writer) Bytes() []byte {
	return w.buf
}
