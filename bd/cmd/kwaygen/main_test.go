package main

import (
	"bytes"
	"strconv"
	"testing"
)

type DevNullStringWriter struct{}

func (DevNullStringWriter) WriteString(s string) (int, error) {
	return len(s), nil
}

func BenchmarkWriteRandNumber(b *testing.B) {
	var buf DevNullStringWriter
	for i := 0; i < b.N; i++ {
		writeRandNumber(42, buf)
	}
}

func TestWriteRandNumber(t *testing.T) {
	var buf bytes.Buffer
	if err := writeRandNumber(42, &buf); err != nil {
		t.Fatal(err)
	}
	bts := buf.Bytes()
	if len(bts) == 0 {
		t.Fatalf("no bytes written")
	}
	if act, exp := bytes.Count(bts, []byte{'\n'}), 1; act != exp {
		t.Fatalf(
			"unexpected number of new lines: %v; want %v",
			act, exp,
		)
	}
	if last := bts[len(bts)-1]; last != '\n' {
		t.Fatalf("new line at unexpected position")
	}

	bts = bts[:len(bts)-1]

	act, err := strconv.Atoi(string(bts))
	if err != nil {
		t.Fatal(err)
	}
	if exp := 42; act != exp {
		t.Fatalf("unexpected number written: %v; want %v", act, exp)
	}
}