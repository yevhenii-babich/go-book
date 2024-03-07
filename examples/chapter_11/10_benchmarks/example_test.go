package main

import "testing"

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
	b.Log("BenchmarkAdd iterations:", b.N)
}
