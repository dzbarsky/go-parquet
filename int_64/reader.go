package int_64

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

func (r *Reader) Next() int64 {
	if r.err != nil {
		return 0
	}

	currOffset := r.offset
	r.offset += 8
	if r.offset > len(r.data) {
		r.err = io.EOF
		return 0
	}

	return int64(binary.LittleEndian.Uint64(r.data[currOffset:r.offset]))
}

func (r *Reader) Error() error {
	return r.err
}
