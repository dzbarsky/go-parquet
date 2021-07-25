package int_32

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	err    error
	data   []byte
	offset int
}

func NewReader(data []byte) *Reader {
	return &Reader{data: data}
}

func (r *Reader) Next() int32 {
	if r.err != nil {
		return 0
	}

	currOffset := r.offset
	r.offset += 4
	if r.offset > len(r.data) {
		r.err = io.EOF
		return 0
	}

	return int32(binary.LittleEndian.Uint32(r.data[currOffset:r.offset]))
}

func (r *Reader) Error() error {
	return r.err
}
