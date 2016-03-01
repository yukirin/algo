package enum

import "testing"

var data = []int{0, 1, 2, 3, 4, 5, 6, 7}
var f = func(a []int) {}

func BenchmarkPowSet(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		PowSet(data, f)
	}
}

func BenchmarkComb(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Comb(data, 5, f)
	}
}

func BenchmarkRepComb(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RepComb(data, 5, f)
	}
}

func BenchmarkPerm(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Perm(data, 5, f)
	}
}

func BenchmarkRepPerm(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RepPerm(data, 5, f)
	}
}
