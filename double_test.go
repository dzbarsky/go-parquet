package main

import (
	"bytes"
	"context"
	"math/rand"
	"os"
	"testing"
)

type doubleHolder struct {
	D float64 `parquet:"d"`
}

func TestDoubleWriter(t *testing.T) {

	values := []doubleHolder{{1}, {3}, {1}, {6}, {7}, {9}, {3}}

	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		t.Fatal(err)
	}

	data := buf.Bytes()

	os.WriteFile("/tmp/out.parquet", data, 0666)

	f := newFile(data)

	readValues := make([]doubleHolder, f.NumRows())
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

func doubleValues(n int) []doubleHolder {
	r := rand.New(rand.NewSource(0))
	values := make([]doubleHolder, n)
	for i := range values {
		values[i].D = r.Float64()
	}
	return values
}

// BenchmarkDoubleWriterPlain-16    	     952	   1132817 ns/op	 1607599 B/op	      93 allocs/op
func BenchmarkDoubleWriterPlain(b *testing.B) {
	values := doubleValues(100000)
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

// BenchmarkDoubleReaderPlain-16    	    1929	    581847 ns/op	 1607204 B/op	      36 allocs/op
func BenchmarkDoubleReaderPlain(b *testing.B) {
	values := doubleValues(100000)
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
		out := make([]doubleHolder, f.NumRows())
		parse(f, out)
	}
}
