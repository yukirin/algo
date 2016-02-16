package astar

import (
	"container/heap"
	"math"
)

// Node is node
type Node struct {
	h    float64
	Cost float64

	Pos  [2]int
	Prev *Node
}

// Q is priority queue
type Q []*Node

func (q Q) Len() int {
	return len(q)
}

func (q Q) Less(i, j int) bool {
	return q[i].h+q[i].Cost < q[j].h+q[j].Cost
}

func (q Q) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push is enqueue
func (q *Q) Push(x interface{}) {
	*q = append(*q, x.(*Node))
}

// Pop is dequeue
func (q *Q) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[0 : len(*q)-1]
	return x
}

// Search is A* algorithm
func Search(maze [][]int, s, g [2]int) *Node {
	move := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	node := &Node{Pos: s}

	maze[s[1]][s[0]] = 1

	q := make(Q, 0, len(maze))
	heap.Init(&q)
	heap.Push(&q, node)

	for q.Len() > 0 {
		n := heap.Pop(&q).(*Node)
		if n.Pos == g {
			return n
		}

		for _, v := range move {
			dst := [2]int{n.Pos[0] + v[0], n.Pos[1] + v[1]}
			if maze[dst[1]][dst[0]] == 1 {
				continue
			}

			forecast := math.Hypot(float64(g[0]-dst[0]), float64(g[1]-dst[1]))
			next := &Node{
				h:    forecast,
				Cost: n.Cost + 1,
				Pos:  dst,
				Prev: n,
			}

			heap.Push(&q, next)
			maze[dst[1]][dst[0]] = 1
		}
	}
	return nil
}
