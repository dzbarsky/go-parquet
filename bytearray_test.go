package main

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"os"
	"testing"
)

type bytearrayHolder struct {
	B []byte `parquet:"b"`
}

func TestBytearrayWriter(t *testing.T) {
	values := []bytearrayHolder{
		{B: []byte{}},
		{B: []byte("HELLO")},
		{B: []byte("WORLD")},
	}

	testCases := map[string][]writeOption{
		"Default": {},
	}

	for name, opts := range testCases {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			err := write(context.Background(), buf, values, opts...)
			if err != nil {
				t.Fatal(err)
			}

			data := buf.Bytes()

			os.WriteFile("/tmp/out.parquet", data, 0666)

			f := newFile(data)

			readValues := make([]bytearrayHolder, f.NumRows())
			parse(f, readValues)

			if len(values) != len(readValues) {
				t.Fatal("bad length")
			}

			for i := range values {
				if !bytes.Equal(values[i].B, readValues[i].B) {
					t.Fatalf("Bad at index %v, wanted %v, got %v", i, values[i], readValues[i])
				}
			}
		})
	}
}

func bytearrayValues(n int) []bytearrayHolder {
	r := rand.New(rand.NewSource(0))
	values := make([]bytearrayHolder, n)
	for i := range values {
		for j := 0; j < r.Intn(100); j++ {
			values[i].B = append(values[i].B, byte(r.Intn(255)))
		}
	}
	return values
}

// BenchmarkBytearrayWriterPlain-16    	     796	   1428661 ns/op	 1525556 B/op	      90 allocs/op
func BenchmarkBytearrayWriterPlain(b *testing.B) {
	values := bytearrayValues(100000)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := write(context.Background(), io.Discard, values)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkBytearrayReaderPlain-16    	     277	   4340098 ns/op	 5291546 B/op	   99010 allocs/op
func BenchmarkBytearrayReaderPlain(b *testing.B) {
	values := bytearrayValues(100000)
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
		out := make([]bytearrayHolder, f.NumRows())
		parse(f, out)
	}
}
