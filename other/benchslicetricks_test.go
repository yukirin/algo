package main

import "testing"

var f = func([]int) {}

func BenchmarkCopy(b *testing.B) {
	a := make([]int, 100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := make([]int, 100000)
		copy(buf, a)
		f(buf)
	}
}

func BenchmarkAppend(b *testing.B) {
	a := make([]int, 100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := append([]int(nil), a...)
		f(buf)
	}
}

func BenchmarkAppendFor(b *testing.B) {
	buf := make([]int, 100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := make([]int, 10)
		for _, v := range buf {
			a = append(a, v)
		}
		f(a)
	}
}

func BenchmarkAppendDot(b *testing.B) {
	buf := make([]int, 100000)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := make([]int, 10)
		a = append(a, buf...)
		f(a)
	}
}
