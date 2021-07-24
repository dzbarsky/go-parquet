package main

import (
	"bytes"
	"context"
	"os"
	"testing"
)

func TestFloatWriter(t *testing.T) {
	type v struct {
		V float32 `parquet:"v"`
	}

	values := []v{{1}, {3}, {1}, {6}, {7}, {9}, {3}}

	buf := bytes.NewBuffer(nil)
	err := write(context.Background(), buf, values)
	if err != nil {
		t.Fatal(err)
	}

	data := buf.Bytes()

	os.WriteFile("/tmp/out.parquet", data, 0666)

	f := newFile(data)

	readValues := make([]v, f.NumRows())
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
