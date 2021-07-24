package bytearray

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	err error
	data []byte
	offset int
}

func NewReader(data []byte) *Reader {
	return &Reader{data: data}
}

func (r *Reader) Next() []byte {
	if r.err != nil {
		return nil
	}

	currOffset := r.offset
	r.offset += 4
	if r.offset > len(r.data) {
		r.err = io.EOF
		return nil
	}

	n := binary.LittleEndian.Uint32(r.data[currOffset:r.offset])

	currOffset = r.offset
	r.offset += int(n)
	if r.offset > len(r.data) {
		r.err = io.EOF
		return nil
	}

	ret := make([]byte, n)
	copy(ret, r.data[currOffset:r.offset])
	return ret
}

func (r *Reader) Error() error {
	return r.err
}
