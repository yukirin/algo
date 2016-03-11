package dijkstra

import (
	"container/heap"
	"math"
)

type priorityQ [][]int

func (q priorityQ) Len() int {
	return len(q)
}

func (q priorityQ) Less(i, j int) bool {
	return q[i][1] < q[j][1]
}

func (q priorityQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *priorityQ) Push(x interface{}) {
	*q = append(*q, x.([]int))
}

func (q *priorityQ) Pop() interface{} {
	x := (*q)[len(*q)-1]
	*q = (*q)[0 : len(*q)-1]
	return x
}

//Search is dijkstra's algorithm
func Search(s int, adjList [][][]int) []int {
	count := 0
	used := make([]bool, len(adjList))
	dist := make([]int, len(adjList))
	for i := range dist {
		dist[i] = math.MaxInt64
	}

	q := make(priorityQ, 0, len(adjList))
	heap.Init(&q)
	heap.Push(&q, []int{s, 0})

	for q.Len() != 0 {
		n := q.Pop().([]int)
		if used[n[0]] {
			continue
		}

		used[n[0]] = true
		dist[n[0]] = n[1]
		count++
		if count == len(adjList) {
			break
		}

		for _, next := range adjList[n[0]] {
			cost := n[1] + next[1]
			if cost >= dist[next[0]] {
				continue
			}
			dist[next[0]] = cost
			heap.Push(&q, []int{next[0], cost})
		}
	}
	return dist
}
