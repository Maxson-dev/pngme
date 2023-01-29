package utils

import (
	"bytes"
	"testing"
)

func TestB2S(t *testing.T) {
	type TestCase struct {
		In  []byte
		Out string
	}

	strs := []string{
		"hello",
		"world",
		"ascii",
		"test",
		"9kpZi1IjwylKjgKmY2jg",
		"bx8evCBl8eM59NewUUmy",
		"nt4k1fVSXtMx4OtYFuKN",
		"русские 234324буквы",
		"康熙字234典體2232testing",
	}

	tests := make([]TestCase, 0, len(strs))

	for i := range strs {
		in := make([]byte, 0, len(strs[i]))
		in = append(in, strs[i]...)
		tests = append(tests, TestCase{In: in, Out: strs[i]})
	}
	for _, test := range tests {
		res := B2S(test.In)
		t.Logf("B2S(%v) == %s\n", test.In, test.Out)
		if res != test.Out {
			t.Fatalf("B2S(%v) == %s\n", test.In, test.Out)
		}
	}
}

func TestS2B(t *testing.T) {
	type TestCase struct {
		In  string
		Out []byte
	}

	strs := []string{
		"hello",
		"world",
		"ascii",
		"test",
		"9kpZi1IjwylKjgKmY2jg",
		"bx8evCBl8eM59NewUUmy",
		"nt4k1fVSXtMx4OtYFuKN",
		"русские 234324буквы",
		"康熙字234典體2232testing",
	}

	tests := make([]TestCase, 0, len(strs))

	for _, in := range strs {
		out := make([]byte, 0, len(in))
		out = append(out, in...)
		tests = append(tests, TestCase{In: in, Out: out})
	}

	for _, test := range tests {
		res := S2B(test.In)
		t.Logf("S2B(%s) == %v\n", test.In, test.Out)
		if !bytes.Equal(res, test.Out) {
			t.Fatalf("S2B(%s) == %v\n", test.In, test.Out)
		}
	}
}
