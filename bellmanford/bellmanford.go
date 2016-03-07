package bellmanford

import "fmt"

// Solve solves the shortest path of the single-source
func Solve(start int, vers int, edges [][]int) ([]int, error) {
	d := make([]int, vers)
	for i := range d {
		d[i] = 1e15
	}
	d[start] = 0

	vers--
	for i := 0; i < vers; i++ {
		for _, v := range edges {
			from, to, cost := v[0], v[1], v[2]
			if d[from]+cost >= d[to] {
				continue
			}
			d[to] = d[from] + cost
		}
	}

	for _, v := range edges {
		from, to, cost := v[0], v[1], v[2]
		if d[from]+cost < d[to] {
			return nil, fmt.Errorf("negative cycle")
		}
	}
	return d, nil
}
