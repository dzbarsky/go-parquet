package main

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"os"
	"testing"
)

type int32Holder struct {
	I int32 `parquet:"i"`
}

func TestInt32Writer(t *testing.T) {

	values := []int32Holder{{1}, {3}, {1}, {6}, {7}, {9}, {3}}

	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		t.Fatal(err)
	}

	data := buf.Bytes()

	os.WriteFile("/tmp/out.parquet", data, 0666)

	f := newFile(data)

	readValues := make([]int32Holder, f.NumRows())
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

func int32Values(n int) []int32Holder {
	r := rand.New(rand.NewSource(0))
	values := make([]int32Holder, n)
	for i := range values {
		values[i].I = r.Int31()
	}
	return values
}

// BenchmarkInt32WriterPlain-16    	    4335	    241889 ns/op	  403192 B/op	      88 allocs/op
func BenchmarkInt32WriterPlain(b *testing.B) {
	values := int32Values(100000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := write(context.Background(), io.Discard, values)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkInt32ReaderPlain-16    	    3607	    309951 ns/op	  804386 B/op	      36 allocs/op
func BenchmarkInt32ReaderPlain(b *testing.B) {
	values := int32Values(100000)
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
		out := make([]int32Holder, f.NumRows())
		parse(f, out)
	}
}
