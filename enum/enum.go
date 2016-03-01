package enum

// PowSet enumerate power set
func PowSet(a []int, fn func([]int)) {
	l := uint(len(a))
	max, ret := 1<<l, make([]int, 0, len(a))

	for i := 0; i < max; i++ {
		n := i
		for j := uint(0); j < l; j++ {
			if n&1 == 1 {
				ret = append(ret, a[j])
			}
			n >>= 1
		}
		fn(ret)
		ret = ret[:0]
	}
}

// Comb enumerate the combination
func Comb(a []int, r int, fn func([]int)) {
	if r > len(a) || r < 1 {
		return
	}

	var f func([]int, int, int)
	f = func(ret []int, last, c int) {
		if c == 0 {
			fn(ret)
			return
		}

		c--
		for i, v := range a[last : len(a)-c] {
			ret[r-c-1] = v
			f(ret, i+last+1, c)
		}
	}

	ret := make([]int, r)
	f(ret, 0, r)
}

// RepComb enumerate the repeated combination
func RepComb(a []int, r int, fn func([]int)) {
	if r < 1 {
		return
	}

	var f func([]int, int, int)
	f = func(ret []int, last, c int) {
		if c == 0 {
			fn(ret)
			return
		}

		for i, v := range a[last:] {
			ret[r-c] = v
			f(ret, i+last, c-1)
		}
	}

	ret := make([]int, r)
	f(ret, 0, r)
}

// Perm enumerate the permutation
// https://oku.edu.mie-u.ac.jp/~okumura/algo/genperm.c
func Perm(a []int, r int, fn func([]int)) {
	permL := make([]int, r)
	f := func(s []int) {
		copy(permL, s)

		k, i := 1, 0
		c := make([]int, r+1)
		for i := 1; i < len(c); i++ {
			c[i] = i
		}

		for k < r {
			if k&1 == 1 {
				i = c[k]
			} else {
				i = 0
			}

			permL[k], permL[i] = permL[i], permL[k]

			fn(permL)
			k = 1
			for ; c[k] == 0; k++ {
				c[k] = k
			}
			c[k]--
		}
	}

	if r < 2 {
		Comb(a, r, fn)
		return
	}
	Comb(a, r, f)
}

// RepPerm enumerae the repeated permutation
func RepPerm(a []int, r int, fn func([]int)) {
	ret := make([]int, r)
	var f func(c int)
	f = func(c int) {
		if c == 0 {
			fn(ret)
			return
		}

		for _, v := range a {
			ret[r-c] = v
			f(c - 1)
		}
	}
	f(r)
}
