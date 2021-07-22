package main

import (
	"os"
	"math"
	"testing"
	"strconv"
)

func TestByteArray(t *testing.T) {
	data, err := os.ReadFile("testdata/pandas/tiny_byte_array.parquet")
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
	data, err := os.ReadFile("testdata/pandas/tiny_float.parquet")
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

func f64Equal(f1, f2 float64) bool {
	if f1 == f2 {
		return true
	}
	return math.IsNaN(f1) && math.IsNaN(f2)
}

func TestFloatWithNaN(t *testing.T) {
	data, err := os.ReadFile("testdata/pandas/float_with_nan.parquet")
	if err != nil {
		t.Fatal(err)
	}
	structs := parse(data)
	if len(structs) != 4 {
		t.Fatal("Wrong length")
	}
	expected := []float64{1, math.NaN(), 2, math.NaN()}
	for i, s := range structs {
		if !f64Equal(s.TestFloat, expected[i]) {
			t.Fatalf("Wrong at %d, wanted %v, got %v", i, expected[i], s.TestFloat)
		}
	}
}


func TestInt(t *testing.T) {
	for _, f := range []string {"testdata/pandas/tiny_int.parquet", "testdata/pandas/tiny_int.snappy.parquet"} {
		t.Run(f, func (t *testing.T) {
			data, err := os.ReadFile(f)
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
		})
	}
}
