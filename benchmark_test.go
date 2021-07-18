package main

import (
	"os"
	"testing"
)

var out []Test

// BenchmarkParse-16    	     182	   6275511 ns/op	11681615 B/op	     194 allocs/op
// BenchmarkParse-16    	     340	   3388624 ns/op	 2434344 B/op	     166 allocs/op
// BenchmarkParse-16    	     732	   1553712 ns/op	 2434323 B/op	     166 allocs/op
func BenchmarkParseRepeated(b *testing.B) {
	benchmarkParse(b, "repeated_bigger.parquet")
}

// BenchmarkParseRandom-16    	    1098	   1093103 ns/op	  833488 B/op	   41253 allocs/op
func BenchmarkParseRandom(b *testing.B) {
	benchmarkParse(b, "random.parquet")
}

func benchmarkParse(b *testing.B, filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out = parse(data)
	}
}
