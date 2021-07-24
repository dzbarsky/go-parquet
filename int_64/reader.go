package int_64

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Reader struct {
	r   io.Reader
	err error
	buf [8]byte
}

func NewReader(data []byte) *Reader {
	return &Reader{r: bytes.NewReader(data)}
}

func (r *Reader) Next() int64 {
	if r.err != nil {
		return 0
	}
	data := r.buf[:]
	// TODO: io.ReadFull would be safer
	_, r.err = r.r.Read(data)
	return int64(binary.LittleEndian.Uint64(data))
}

func (r *Reader) Error() error {
	return r.err
}
