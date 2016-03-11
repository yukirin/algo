// Package segment is (lazy) segment tree
// http://www.npca.jp/works/magazine/2015_5/
package segment

import "math"

// NewLazy return the initialized lazy segment tree (n == 2^p)
func NewLazy(n int) ([]int64, []int64) {
	t := New(n)
	return t, make([]int64, len(t))
}

// NewLazyS returns the initialization lazy segment tree value of the slice
func NewLazyS(inits []int64) ([]int64, []int64) {
	t := NewS(inits)
	return t, make([]int64, len(t))
}

// LazyUpdate update the lazy segment tree
func LazyUpdate(a, b, l, r, i int, x int64, t, lazy []int64) {
	if r <= a || b <= l {
		return
	}

	if a <= l && r <= b {
		lazy[i] += x
		lazyEval(i, t, lazy)
		return
	}

	lazyEval(i, t, lazy)
	LazyUpdate(a, b, l, (l+r)/2, i*2+1, x, t, lazy)
	LazyUpdate(a, b, (l+r)/2, r, i*2+2, x, t, lazy)

	if t[i*2+1] < t[i*2+2] {
		t[i] = t[i*2+1]
	} else {
		t[i] = t[i*2+2]
	}
	return
}

// LazyQuery to run a query in the interval
func LazyQuery(a, b, l, r, i int, t, lazy []int64) int64 {
	lazyEval(i, t, lazy)

	if r <= a || b <= l {
		return math.MaxInt64
	}

	if a <= l && r <= b {
		return t[i]
	}

	vl := LazyQuery(a, b, l, (l+r)/2, i*2+1, t, lazy)
	vr := LazyQuery(a, b, (l+r)/2, r, i*2+2, t, lazy)

	if vl < vr {
		return vl
	}
	return vr
}

func lazyEval(i int, t, lazy []int64) {
	t[i] += lazy[i]
	if i < (len(t)+1)/2-1 {
		lazy[i*2+1] += lazy[i]
		lazy[i*2+2] += lazy[i]
	}
	lazy[i] = 0
}
