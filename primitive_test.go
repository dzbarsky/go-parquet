package main

import (
	"os"
	"testing"
	"strconv"
)

func TestByteArray(t *testing.T) {
	data, err := os.ReadFile("testdata/tiny_byte_array.parquet")
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

func TestFloat(t *testing.T) {
	data, err := os.ReadFile("testdata/tiny_float.parquet")
	if err != nil {
		t.Fatal(err)
	}
	structs := parse(data)
	if len(structs) != 10 {
		t.Fatal("Wrong length")
	}
	for i, s := range structs {
		if s.TestFloat != float64(i) {
			t.Fatalf("Wrong at %d", i)
		}
	}
}

func TestInt(t *testing.T) {
	data, err := os.ReadFile("testdata/tiny_int.parquet")
	if err != nil {
		t.Fatal(err)
	}
	structs := parse(data)
	if len(structs) != 10 {
		t.Fatal("Wrong length")
	}
	for i, s := range structs {
		if s.TestInt != int64(i) {
			t.Fatalf("Wrong at %d", i)
		}
	}
}
