package main

import (
	"os"
	"testing"
)

var out []Test

// BenchmarkParse-16    	     182	   6275511 ns/op	11681615 B/op	     194 allocs/op
// BenchmarkParse-16    	     340	   3388624 ns/op	 2434344 B/op	     166 allocs/op
// BenchmarkParse-16    	     732	   1553712 ns/op	 2434323 B/op	     166 allocs/op
func BenchmarkParse(b *testing.B) {
	data, err := os.ReadFile("repeated_bigger.parquet")
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out = parse(data)
	}
}
