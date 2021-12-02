package main

import "testing"

func BenchmarkVersion1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Version1()
	}
}
