package main

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"os"
	"testing"
)

type int64Holder struct {
	I int64 `parquet:"i"`
}

func TestInt64Writer(t *testing.T) {

	values := []int64Holder{{1}, {3}, {1}, {6}, {7}, {9}, {3}}

	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		t.Fatal(err)
	}

	data := buf.Bytes()

	os.WriteFile("/tmp/out.parquet", data, 0666)

	f := newFile(data)

	readValues := make([]int64Holder, f.NumRows())
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

func int64Values(n int) []int64Holder {
	r := rand.New(rand.NewSource(0))
	values := make([]int64Holder, n)
	for i := range values {
		values[i].I = r.Int63()
	}
	return values
}

// BenchmarkInt64WriterPlain-16    	    2845	    368690 ns/op	  804592 B/op	      88 allocs/op
func BenchmarkInt64WriterPlain(b *testing.B) {
	values := int64Values(100000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := write(context.Background(), io.Discard, values)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkInt64ReaderPlain-16    	    2209	    536778 ns/op	 1607200 B/op	      36 allocs/op
func BenchmarkInt64ReaderPlain(b *testing.B) {
	values := int64Values(100000)
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
		out := make([]int64Holder, f.NumRows())
		parse(f, out)
	}
}
