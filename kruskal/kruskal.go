package kruskal

import (
	"container/heap"

	"github.com/yukirin/algo/unionfind"
)

// Edge is graph edge
type Edge struct {
	Ps     []int
	Weight int
}

type priorityQ []Edge

func (h priorityQ) Len() int {
	return len(h)
}

func (h priorityQ) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

func (h priorityQ) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *priorityQ) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *priorityQ) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

// Solve is クラスカル法で最小全域木を求める
func Solve(vertices int, edges []Edge) (int, []Edge) {
	var t unionfind.Tree
	var (
		es   []Edge
		dist int
	)

	t.Init(vertices)

	q := &priorityQ{}
	heap.Init(q)

	for _, e := range edges {
		heap.Push(q, e)
	}

	for q.Len() > 0 {
		e := heap.Pop(q).(Edge)
		if t.Same(e.Ps[0], e.Ps[1]) {
			continue
		}

		t.Union(e.Ps[0], e.Ps[1])
		es = append(es, e)
		dist += e.Weight
	}

	return dist, es
}
