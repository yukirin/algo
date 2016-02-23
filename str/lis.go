package str

import "sort"

// LIS is longest increasing subsequence
func LIS(seq []int) []int {
	inf := 1<<63 - 1
	max := 0

	ret := make([]int, len(seq))
	dp := make([]int, len(seq))
	for i := range dp {
		dp[i] = inf
		ret[i] = inf
	}

	for _, n := range seq {
		i := sort.SearchInts(dp, n)
		dp[i] = n
		if i+1 > max {
			max = i + 1
			copy(ret, dp[:i+1])
		}
	}

	return ret[:max]
}
