package segment

// New return the initialized segment tree Inf (n == 2^p)
func New(n int) []int64 {
	if n&(n-1) != 0 {
		return nil
	}

	const inf = 2<<62 - 1

	t := make([]int64, n*2-1)
	for i := range t {
		t[i] = inf
	}
	return t
}

// NewS returns the initialization segment tree value of the slice (len(inits) == 2^p)
func NewS(n int, inits []int64) []int64 {
	if len(inits)&(len(inits)-1) != 0 {
		return nil
	}

	t := make([]int64, n*2-1)
	for i := 0; i < n; i++ {
		t[n-1+i] = inits[i]
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
func Query(a, b, i, l, r int, t []int64) int64 {
	const inf = 2<<62 - 1
	if r <= a || b <= l {
		return inf
	}

	if a <= l && r <= b {
		return t[i]
	}

	vl := Query(a, b, i*2+1, l, (l+r)/2, t)
	vr := Query(a, b, i*2+2, (l+r)/2, r, t)

	if vl < vr {
		return vl
	}
	return vr
}
