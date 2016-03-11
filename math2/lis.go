package math2

import (
	"math"
	"sort"
)

// LIS is longest increasing subsequence
func LIS(seq []int) int {
	dp := make([]int, len(seq))
	for i := range dp {
		dp[i] = math.MaxInt64
	}

	for _, n := range seq {
		i := sort.SearchInts(dp, n)
		dp[i] = n
	}

	return sort.SearchInts(dp, math.MaxInt64)
}
