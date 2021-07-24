package main

import (
	"bytes"
	"context"
	"math/rand"
	"os"
	"testing"
)

type floatHolder struct {
	F float32 `parquet:"f"`
}

func TestFloatWriter(t *testing.T) {

	values := []floatHolder{{1}, {3}, {1}, {6}, {7}, {9}, {3}}

	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		t.Fatal(err)
	}

	data := buf.Bytes()

	os.WriteFile("/tmp/out.parquet", data, 0666)

	f := newFile(data)

	readValues := make([]floatHolder, f.NumRows())
	parse(f, readValues)

	if len(values) != len(readValues) {
		t.Fatal("bad length")
	}

	for i := range values {
		if values[i] != readValues[i] {
			t.Fatal("Bad at index", i)
		}
	}
}

func floatValues(n int) []floatHolder {
	r := rand.New(rand.NewSource(0))
	values := make([]floatHolder, n)
	for i := range values {
		values[i].F = r.Float32()
	}
	return values
}

var data []byte

// BenchmarkFloatWriterPlain-16    	    1030	   1131445 ns/op	 1530206 B/op	     106 allocs/op
func BenchmarkFloatWriterPlain(b *testing.B) {
	values := floatValues(100000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(nil)
		err := write(context.Background(), buf, values)
		if err != nil {
			b.Fatal(err)
		}

		data = buf.Bytes()
	}
}

// BenchmarkFloatReaderPlain-16    	    3846	    318899 ns/op	  804387 B/op	      36 allocs/op
func BenchmarkFloatReaderPlain(b *testing.B) {
	values := floatValues(100000)
	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		b.Fatal(err)
	}
	data = buf.Bytes()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		f := newFile(data)
		out := make([]floatHolder, f.NumRows())
		parse(f, out)
	}
}
