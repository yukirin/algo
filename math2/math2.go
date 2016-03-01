package math2

import (
	"math"
	"time"
)

// Epsilon is machine epsilon
var Epsilon = math.Nextafter(1, 2) - 1

// Factor is prime factorization
func Factor(n uint64) []uint64 {
	if n <= 1 {
		return []uint64{n}
	}

	ps := make([]uint64, 0, 100)

	for i := uint64(2); i*i <= n; i++ {
		for n%i == 0 {
			n /= i
			ps = append(ps, i)
		}
	}
	if n > 1 {
		ps = append(ps, n)
	}
	return ps
}

// Gcd is greatest common divisor
func Gcd(a, b uint64) uint64 {
	for ; b != 0; b, a = a%b, b {
	}
	return a
}

// ExGcd is Extended Euclidean algorithm
func ExGcd(x, y int64) (a, b, c int64) {
	if x <= 0 || y <= 0 {
		return
	}

	r0, r1 := x, y
	a0, a1 := int64(1), int64(0)
	b0, b1 := int64(0), int64(1)

	for r1 > 0 {
		q1, r2 := r0/r1, r0%r1
		a2, b2 := a0-q1*a1, b0-q1*b1
		r0, r1 = r1, r2
		a0, a1 = a1, a2
		b0, b1 = b1, b2
	}
	a, b, c = a0, b0, r0
	return
}

// ChineseRT is chinese remainder theorem
func ChineseRT(a1, m1, a2, m2 int64) int64 {
	a, _, c := ExGcd(m1, m2)
	if c != 1 {
		return -1
	}

	x := a1 + (a2-a1)*a*m1
	for x < 0 {
		x += m1 * m2
	}
	return x
}

// Lcm is least common multiple
func Lcm(a, b uint64) uint64 {
	return (a * b) / Gcd(a, b)
}

// Round is rounding
func Round(f float64, place int) float64 {
	shift := math.Pow10(place)
	return math.Floor(f*shift+.5) / shift
}

// Round1 is rounded to the nearest decimal one
func Round1(f float64) float64 {
	return math.Floor(f + .5)
}

// BitCount is count the number of bits where 1 is set
// http://qiita.com/fmhr/items/fa5a7d9b785456446768
func BitCount(x uint64) uint64 {
	x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
	x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)
	x = (x & 0x00ff00ff00ff00ff) + ((x & 0xff00ff00ff00ff00) >> 8)
	x = (x & 0x0000ffff0000ffff) + ((x & 0xffff0000ffff0000) >> 16)
	x = (x & 0x00000000ffffffff) + ((x & 0xffffffff00000000) >> 32)
	return x
}

// PowExp is 二分累乗法
func PowExp(n, e uint64) uint64 {
	a := uint64(1)
	for e > 0 {
		if e&1 == 1 {
			a *= n
		}
		n *= n
		e >>= 1
	}
	return a
}

// PowMod is 繰り返し自乗法
func PowMod(n, e, mod uint64) uint64 {
	a := uint64(1)
	for e > 0 {
		if e&1 == 1 {
			a *= n
			a %= mod
		}
		n *= n
		n %= mod
		e >>= 1
	}
	return a % mod
}

// Xor128 is pseudo random number generator
var Xor128 func() uint64

func init() {
	seed := []uint64{123456789, 362436069, 521288629, uint64(time.Now().UnixNano())}
	Xor128 = func() uint64 {
		t := seed[0] ^ (seed[0] << 11)
		seed[0], seed[1], seed[2] = seed[1], seed[2], seed[3]
		seed[3] = (seed[3] ^ (seed[3] >> 19)) ^ (t ^ (t >> 8))
		return seed[3]
	}
}
