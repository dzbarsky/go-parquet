package bytearray

import (
	"io"
	"encoding/binary"
)

type reader interface {
	io.Reader
	io.ByteReader
}

type Reader struct{
	r reader
	err error
	buf [4]byte
}

func NewReader(r reader) *Reader {
	return &Reader{r: r}
}

func (r* Reader) Next() []byte {
	if r.err != nil {
		return nil
	}

	data := r.buf[:]
	_, r.err = r.r.Read(data)
	if r.err != nil {
		return nil
	}

	n := binary.LittleEndian.Uint32(data)

	data = make([]byte, n)
	_, r.err = io.ReadFull(r.r, data)
	return data
}

func (r* Reader) Error() error {
	return r.err
}
