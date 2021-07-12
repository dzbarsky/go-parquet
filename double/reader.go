package double

import (
	"io"
	"math"
	"encoding/binary"
)

type Reader struct{
	r io.Reader
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r* Reader) Next() float64 {
	if r.err != nil {
		return 0
	}
	var data uint64
	r.err = binary.Read(r.r, binary.LittleEndian, &data)
	return math.Float64frombits(data)
}

func (r* Reader) Error() error {
	return r.err
}
