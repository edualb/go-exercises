package main

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3([]string{"o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o", "o"})
	}
}
