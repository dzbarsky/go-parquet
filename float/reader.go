package float

import (
	"encoding/binary"
	"io"
	"math"
)

type Reader struct {
	err    error
	data   []byte
	offset int
}

func NewReader(data []byte) *Reader {
	return &Reader{data: data}
}

func (r *Reader) Next() float32 {
	if r.err != nil {
		return 0
	}

	currOffset := r.offset
	r.offset += 4
	if r.offset > len(r.data) {
		r.err = io.EOF
		return 0
	}

	return math.Float32frombits(binary.LittleEndian.Uint32(r.data[currOffset:r.offset]))
}

func (r *Reader) Error() error {
	return r.err
}
