package main

import (
	"os"
	"testing"
)

var out []Test

// BenchmarkParse-16    	     182	   6275511 ns/op	11681615 B/op	     194 allocs/op
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
