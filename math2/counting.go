package math2

import (
	"math/big"

	"github.com/yukirin/algo/sieve"
)

// Fact is n!
func Fact(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	ans := uint64(1)
	for i := uint64(1); i <= n; i++ {
		ans *= i
	}
	return ans
}

// FactM is n! % p
func FactM(n, p uint64) uint64 {
	if n == 0 {
		return 1 % p
	}

	ret := uint64(1)
	for i := uint64(1); i <= n; i++ {
		ret *= i
		ret %= p
	}
	return ret
}

//NPR is nPr
func NPR(n, r uint64) uint64 {
	p := n
	for i := n - r + 1; i < n; i++ {
		p *= i
	}
	return p
}

// NPRM is nPr % p
func NPRM(n, r, p uint64) uint64 {
	ret := n % p
	for i := n - r + 1; i < n; i++ {
		ret *= i
		ret %= p
	}
	return ret
}

//NR is Repeated permutation
func NR(n, r uint64) uint64 {
	return PowExp(n, r)
}

// NRM is Repeated permutation
func NRM(n, r, p uint64) uint64 {
	return PowMod(n, r, p)
}

//NCR is nCr
func NCR(n, r uint64) uint64 {
	if n < 0 || r < 0 || r > n {
		return 0
	}

	k := r
	if k > n/2 {
		k = n - r
	}

	ret := uint64(1)
	for i := uint64(1); i <= k; i++ {
		ret *= n - i + 1
		ret /= i
	}
	return ret
}

// NCRM is nCr % p (r<10^7, n,r < m, m:prime number)
// http://www37.atwiki.jp/uwicoder/pages/2118.html
func NCRM(n, r, p uint64) uint64 {
	ret := uint64(1)
	for {
		if r == 0 {
			break
		}
		n2 := n % p
		r2 := r % p
		if n2 < r2 {
			return 0
		}

		for i := uint64(0); i < r2; i++ {
			ret = ret * (n2 - i) % p
		}

		imul := uint64(1)
		for i := uint64(0); i < r2; i++ {
			imul = imul * (i + 1) % p
		}
		a, b := big.NewInt(int64(imul)), big.NewInt((int64(p)))
		k := uint64(new(big.Int).ModInverse(a, b).Int64())
		ret = ret * k % p
		n /= p
		r /= p
	}
	return ret
}

// NCRM2 is nCr % p (n<10^7, r,m: Any)
// http://www37.atwiki.jp/uwicoder/pages/2118.html
func NCRM2(n, r, mod uint64) uint64 {
	if n < 0 || r < 0 || r > n {
		return 0
	}

	if r > n/2 {
		r = n - r
	}

	a := make([]uint64, n)
	for i := uint64(0); i < n; i++ {
		a[i] = n - i
	}

	ps := sieve.Eratosthenes(int(r))
	for i := range ps {
		p := uint64(ps[i])
		for q := p; q <= r; q *= p {
			m := n % q
			for i, j := m, uint64(0); j < r/q; i, j = i+q, j+1 {
				a[i] /= p
			}
		}
	}

	mul := uint64(1)
	for i := uint64(0); i < r; i++ {
		mul = mul * a[i] % mod
	}
	return mul
}

// NHR is nHr
func NHR(n, r uint64) uint64 {
	return NCR(n+r-1, r)
}

// NHRM is nHr % p
func NHRM(n, r, p uint64) uint64 {
	return NCRM(n+r-1, r, p)
}
