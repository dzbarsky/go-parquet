package main

import (
	"bytes"
	"math/bits"

	"parquet/float"
)

type floatDict struct {
	dictValues []float32
	indices    map[float32]int32

	values []int32
}

func NewFloatDict(nValues int) *floatDict {
	return &floatDict{
		dictValues: make([]float32, 0, nValues),
		indices:    make(map[float32]int32, nValues),
		values:     make([]int32, 0, nValues),
	}
}

func (fd *floatDict) Write(v float32) {
	index, ok := fd.indices[v]
	if !ok {
		index = int32(len(fd.dictValues))
		fd.indices[v] = index
		fd.dictValues = append(fd.dictValues, v)
	}
	fd.values = append(fd.values, index)
}

func (fd *floatDict) NDictValues() int {
	return len(fd.dictValues)
}

func (fd *floatDict) DictBytes() []byte {
	fw := float.NewWriter(len(fd.dictValues))
	for _, v := range fd.dictValues {
		fw.Write(v)
	}
	return fw.Bytes()
}

func (fd *floatDict) DataBytes() (int8, []byte) {
	buf := bytes.NewBuffer(nil)

	bitWidth := bits.Len(uint(len(fd.dictValues) - 1))
	hw := newHybridWriter(buf, bitWidth)

	for _, v := range fd.values {
		err := hw.Write(v)
		must(err)
	}
	must(hw.Flush())
	return int8(bitWidth), buf.Bytes()
}
