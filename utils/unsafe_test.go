package utils

import (
	"testing"
)

func TestStringToBytes(t *testing.T) {
	expect := map[string][]byte{
		"hello": []byte{'h', 'e', 'l', 'l', 'o'},
		"":      []byte{},
	}

	for k, v := range expect {
		if b := StringToBytes(k); len(b) != len(v) {
			t.Fatalf("result is %v expect %v", b, v)
		}

		if b := StringToBytesSimple(k); len(b) != len(v) {
			t.Fatalf("result is %v expect %v", b, v)
		}
	}
}

func TestBytesToString(t *testing.T) {
	expect := map[string][]byte{
		"hello": []byte{'h', 'e', 'l', 'l', 'o'},
		"":      []byte{},
	}

	for k, v := range expect {
		if b := BytesToString(v); b != k {
			t.Fatalf("result is %v expect %v", b, k)
		}

		if b := BytesToStringSimple(v); b != k {
			t.Fatalf("result is %v expect %v", b, k)
		}
	}
}

const N = 50000000

func BenchmarkStringToBytesDirect(b *testing.B) {
	for i := 1; i < N; i++ {
		_ = []byte("hello world!")
	}
}

func BenchmarkStringToBytesUnsafe(b *testing.B) {
	for i := 1; i < N; i++ {
		_ = StringToBytes("hello world!")
	}
}

func BenchmarkBytesToStringDirect(b *testing.B) {
	for i := 1; i < N; i++ {
		_ = string([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	}
}

func BenchmarkBytesToStringUnsafe(b *testing.B) {
	for i := 1; i < N; i++ {
		_ = BytesToString([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	}
}
