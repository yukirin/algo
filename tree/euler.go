package tree

// EulerTour is euler tour ( [begin[i]:end[i]] )
func EulerTour(root int, adjL [][]int) ([]int, []int, []int) {
	tour := make([]int, 0, 2*len(adjL)-1)
	begin, end := make([]int, len(adjL)), make([]int, len(adjL))
	c := 0

	var f func(int)
	f = func(i int) {
		begin[i] = c
		tour = append(tour, i)
		c++
		for _, v := range adjL[i] {
			f(v)
			tour = append(tour, i)
			c++
		}
		end[i] = c
	}

	f(root)
	return tour, begin, end
}
