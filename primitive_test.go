package main

import (
	"os"
	"testing"
	"strconv"
)

func TestByteArray(t *testing.T) {
	data, err := os.ReadFile("testdata/tiny_string.parquet")
	if err != nil {
		t.Fatal(err)
	}
	structs := parse(data)
	if len(structs) != 10 {
		t.Fatal("Wrong length")
	}
	for i, s := range structs {
		if string(s.TestBytes) != strconv.Itoa(i) {
			t.Fatalf("Wrong at %d", i)
		}
	}
}
