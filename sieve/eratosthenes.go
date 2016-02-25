package sieve

import "math"

// Eratosthenes is Sieve of Eratosthenes
func Eratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}

	r := int(math.Floor(math.Sqrt(float64(n))))
	list := make([]bool, n+1)
	list[0], list[1] = true, true

	for i := 2; i <= r; i++ {
		if !list[i] {
			for j := i * i; j <= n; j += i {
				list[j] = true
			}
		}
	}

	l := n / int(math.Ceil(math.Log(float64(n))))
	primes := make([]int, 0, l)
	for i, v := range list {
		if v {
			continue
		}
		primes = append(primes, i)
	}

	return primes
}
