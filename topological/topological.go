package topological

type node struct {
	num     int
	in, out map[int]*node
}

func initNode(vs int, edges [][]int) map[int]*node {
	m := make(map[int]*node)

	for i := 0; i < vs; i++ {
		m[i] = &node{
			num: i,
			in:  make(map[int]*node),
			out: make(map[int]*node),
		}
	}

	for i, edge := range edges {
		for _, v := range edge {
			m[i].out[v] = m[v]
			m[v].in[i] = m[i]
		}
	}
	return m
}

// Sort is topological sort
func Sort(vs int, edges [][]int) []int {
	ns := initNode(vs, edges)
	ret := make([]int, 0, vs)
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
