package math2

// Table is mod table
type Table struct {
	Mod    int64
	Inv    []int64
	Fac    []int64
	FacInv []int64
}

// ModTable create mod table
func ModTable(mod int64) Table {
	n := int64(1000010)
	t := Table{
		Mod:    mod,
		Inv:    make([]int64, n+2),
		Fac:    make([]int64, n+1),
		FacInv: make([]int64, n+1),
	}

	t.Fac[0], t.FacInv[0], t.Inv[1] = 1, 1, 1
	for i := int64(1); i <= n; i++ {
		t.Inv[i+1] = mod - (mod/(i+1))*t.Inv[mod%(i+1)]%mod
		t.Fac[i] = (t.Fac[i-1] * i) % mod
		t.FacInv[i] = (t.FacInv[i-1] * t.Inv[i]) % mod
	}
	return t
}

// FactM is n! % mod
func FactM(n int64, t Table) int64 {
	return t.Fac[n]
}

// NPRM is nPr % mod
func NPRM(n, r int64, t Table) int64 {
	ans := n % t.Mod
	for i := n - r + 1; i < n; i++ {
		ans *= i
		ans %= t.Mod
	}
	return ans
}

// NRM is n**r % mod
func NRM(n, r, mod int64) int64 {
	return int64(PowMod(uint64(n), uint64(r), uint64(mod)))
}

// NCRM is nCr % mod
func NCRM(n, r int64, t Table) int64 {
	if !(0 <= r && r <= n) {
		return 0
	}

	if n < t.Mod {
		return (((t.Fac[n] * t.FacInv[r]) % t.Mod) * t.FacInv[n-r]) % t.Mod
	}

	ret := int64(1)
	for ; r > 0; n, r = n/t.Mod, r/t.Mod {
		n0, r0 := n%t.Mod, r%t.Mod
		if n0 < r0 {
			return 0
		}

		ret = ret * ((((t.Fac[n0] * t.FacInv[r0]) % t.Mod) * t.FacInv[n0-r0]) % t.Mod) % t.Mod
	}
	return ret
}

// NHRM is nHr % mod
func NHRM(n, r int64, t Table) int64 {
	if r == 0 {
		return 1
	}

	return NCRM(n+r-1, r, t)
}
