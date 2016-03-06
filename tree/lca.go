package tree

import "math"

// LCA is Lowest Common Ancestor (Doubling)
func LCA(a, b int, depth []int, par [][]int) int {
	if depth[a] > depth[b] {
		a, b = b, a
	}

	l := len(par[0])
	for i := uint(l - 1); i != 0; i-- {
		if ((depth[b]-depth[a])>>i)&1 == 1 {
			b = par[b][i]
		}
	}

	if a == b {
		return a
	}

	for i := l - 1; i >= 0; i-- {
		if par[a][i] != par[b][i] {
			a, b = par[a][i], par[b][i]
		}
	}
	return par[a][0]
}

// BuildLCA build table
func BuildLCA(root int, adjL [][]int) ([]int, [][]int) {
	par := make([][]int, len(adjL))
	depth := make([]int, len(adjL))

	l := int(math.Ceil(math.Log2(float64(len(adjL))))) + 2
	for i := range par {
		par[i] = make([]int, l)
	}

	var f func(int, int, int)
	f = func(v, p, d int) {
		par[v][0] = p
		depth[v] = d

		for _, v2 := range adjL[v] {
			f(v2, v, d+1)
		}
	}

	f(root, -1, 0)
	fill(par, l)
	return depth, par
}

func fill(par [][]int, l int) {
	for i := 1; i < l; i++ {
		for _, v := range par {
			p1 := v[i-1]
			if p1 == -1 {
				v[i] = -1
				continue
			}
			v[i] = par[p1][i-1]
		}
	}
}
