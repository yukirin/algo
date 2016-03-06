package math2

import "sort"

// LIS is longest increasing subsequence
func LIS(seq []int) int {
	inf := 1<<63 - 1

	dp := make([]int, len(seq))
	for i := range dp {
		dp[i] = inf
	}

	for _, n := range seq {
		i := sort.SearchInts(dp, n)
		dp[i] = n
	}

	return sort.SearchInts(dp, inf)
}
