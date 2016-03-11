package segment

import (
	"math"
)

// New return the initialized segment tree (n == 2^p)
func New(n int) []int64 {
	if n&(n-1) != 0 {
		return nil
	}

	t := make([]int64, n*2-1)
	for i := range t {
		t[i] = math.MaxInt64
	}
	return t
}

// NewS returns the initialization segment tree value of the slice
func NewS(inits []int64) []int64 {
	n := 1
	for n < len(inits) {
		n *= 2
	}

	t := make([]int64, n*2-1)
	for i := 0; i < len(inits); i++ {
		t[n-1+i] = inits[i]
	}

	d := n - len(inits)
	for i := 0; i < d; i++ {
		t[n-1+len(inits)+i] = math.MaxInt64
	}

	for i := n - 2; i >= 0; i-- {
		min, v := t[i*2+1], t[i*2+2]
		if v < min {
			min = v
		}
		t[i] = min
	}
	return t
}

// Update update value
func Update(i int, x int64, t []int64) {
	i += (len(t)+1)/2 - 1
	t[i] = x

	for i > 0 {
		i = (i - 1) / 2
		min, v := t[i*2+1], t[i*2+2]
		if v < min {
			min = v
		}
		t[i] = min
	}
}

// Query is a process for the interval
func Query(a, b, l, r, i int, t []int64) int64 {
	if r <= a || b <= l {
		return math.MaxInt64
	}

	if a <= l && r <= b {
		return t[i]
	}

	vl := Query(a, b, l, (l+r)/2, i*2+1, t)
	vr := Query(a, b, (l+r)/2, r, i*2+2, t)

	if vl < vr {
		return vl
	}
	return vr
}
