package topological

type node struct {
	num     int
	in, out map[int]*node
}

func initNode(vs []int, edges [][]int) map[int]*node {
	m := make(map[int]*node)

	for _, v := range vs {
		m[v] = &node{
			num: v,
			in:  make(map[int]*node),
			out: make(map[int]*node),
		}
	}

	for _, e := range edges {
		m[e[0]].out[e[1]] = m[e[1]]
		m[e[1]].in[e[0]] = m[e[0]]
	}
	return m
}

// Sort is topological sort
func Sort(vs []int, edges [][]int) []int {
	ns := initNode(vs, edges)

	ret := make([]int, 0, len(vs))
	var s []*node

	for _, n := range ns {
		if len(n.in) == 0 {
			s = append(s, n)
		}
	}

	for len(s) != 0 {
		n := s[0]
		s = s[1:]
		ret = append(ret, n.num)

		for _, v := range n.out {
			delete(v.in, n.num)
			if len(v.in) == 0 {
				s = append(s, v)
			}
		}
		n.out = nil
	}

	for _, n := range ns {
		if n.out != nil {
			return nil
		}
	}
	return ret
}
