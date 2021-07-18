package double

import (
	"io"
	"math"
	"encoding/binary"
)

type Reader struct{
	r io.Reader
	err error
	buf [8]byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r* Reader) Next() float64 {
	if r.err != nil {
		return 0
	}
	data := r.buf[:]
	_, r.err = r.r.Read(data)
	return math.Float64frombits(binary.LittleEndian.Uint64(data))
}

func (r* Reader) Error() error {
	return r.err
}
