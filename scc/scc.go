package scc

// Decomposition is Decomposition of SCC
func Decomposition(adjL [][]int) [][]int {
	rev := make([][]int, len(adjL))
	l, mark := make([]int, 0, len(rev)), make([]bool, len(rev))
	var ret [][]int

	var visit func(int)
	visit = func(n int) {
		if mark[n] {
			return
		}

		mark[n] = true
		for _, v := range adjL[n] {
			visit(v)
			rev[v] = append(rev[v], n)
		}
		l = append(l, n)
	}
	visit(0)

	for i := range mark {
		mark[i] = false
	}

	var assign func(int, *[]int)
	assign = func(n int, g *[]int) {
		if mark[n] {
			return
		}

		mark[n] = true
		*g = append(*g, n)
		for _, v := range rev[n] {
			assign(v, g)
		}
	}

	for i := len(l) - 1; i >= 0; i-- {
		g := make([]int, 0, len(rev)/2)
		assign(l[i], &g)
		if len(g) == 0 {
			continue
		}
		ret = append(ret, g)
	}
	return ret
}
