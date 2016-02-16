package knapsack

// Solve solve the knapsack
func Solve(max int, goods [][]int) int {
	dp := make(map[int]int)
	dp[0] = 0

	for _, o := range goods {
		nextDp := make(map[int]int)
		v, w := o[0], o[1]

		for k, v2 := range dp {
			nextDp[k] = v2
			if w+k > max {
				continue
			}

			if v+v2 > nextDp[w+k] {
				nextDp[w+k] = v + v2
			}
		}
		dp = nextDp
	}

	ret := 0
	for _, w := range dp {
		if w > ret {
			ret = w
		}
	}
	return ret
}
