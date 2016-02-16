package edmondskarp

// Solve is Edmonds Karp algorithm
func Solve(s, t int, capa, edges [][]int) (int, [][]int) {
	f := 0
	flow := make([][]int, len(capa))
	for i := range flow {
		flow[i] = make([]int, len(capa))
	}

	for {
		m, p := bfs(s, t, capa, edges, flow)
		if m == 0 {
			break
		}

		f += m
		v := t
		for v != s {
			u := p[v]
			flow[u][v] += m
			flow[v][u] -= m
			v = u
		}
	}
	return f, flow
}

func bfs(s, t int, capa, edges, flow [][]int) (int, []int) {
	min := 1<<32 - 1
	p := make([]int, len(edges))
	for i := range p {
		p[i] = -1
	}
	p[s] = -2

	q := make([]int, 0, len(edges))
	q = append(q, s)

	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		for _, v := range edges[u] {
			rest := capa[u][v] - flow[u][v]
			if rest <= 0 || p[v] != -1 {
				continue
			}

			p[v] = u
			if rest < min {
				min = rest
			}

			if v == t {
				return min, p
			}
			q = append(q, v)
		}
	}
	return 0, p
}
