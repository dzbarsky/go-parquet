package main

import (
	"os"
	"testing"
)

type s struct {
	Test      float64 `parquet:"test"`
	TestFloat float64 `parquet:"test_float"`
	TestInt   int64   `parquet:"test_int"`
	TestBytes []byte  `parquet:"test_str"`
}

// BenchmarkParseRepeated-16    	    1096	    968192 ns/op	  828591 B/op	     133 allocs/op
func BenchmarkParseRepeated(b *testing.B) {
	benchmarkParse(b, "scratch/repeated_bigger.parquet")
}

// BenchmarkParseRandom-16    	    3637	    275558 ns/op	  207894 B/op	     129 allocs/op
func BenchmarkParseRandom(b *testing.B) {
	benchmarkParse(b, "scratch/random.parquet")
}

func benchmarkParse(b *testing.B, filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	f := newFile(data)
	out := make([]s, f.NumRows())

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		f = newFile(data)
		parse(f, out)
	}
}
