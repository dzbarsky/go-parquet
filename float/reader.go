package float

import (
	"io"
	"math"
	"encoding/binary"
)

type Reader struct{
	r io.Reader
	err error
	buf [4]byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r* Reader) Next() float32 {
	if r.err != nil {
		return 0
	}
	data := r.buf[:]
	// TODO: io.ReadFull would be safer
	_, r.err = r.r.Read(data)
	return math.Float32frombits(binary.LittleEndian.Uint32(data))
}

func (r* Reader) Error() error {
	return r.err
}
