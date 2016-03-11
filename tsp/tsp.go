package tsp

import (
	"math"
)

// Solve solve the TSP
func Solve(towns int, dist [][]float64) float64 {
	n := uint(towns)
	dp := make([][]float64, 1<<n)

	for i := 0; i < 1<<n; i++ {
		for j := uint(0); j < n; j++ {
			dp[i] = append(dp[i], math.MaxFloat64)
		}
	}
	dp[1][0] = 0

	for i := 0; i < 1<<n; i++ {
		for j := uint(0); j < n; j++ {
			if dp[i][j] == math.MaxFloat64 {
				continue
			}

			for k := uint(0); k < n; k++ {
				if (i>>k)&1 == 1 {
					continue
				}

				nextI := i | 1<<k
				nextD := dp[i][j] + dist[j][k]
				if nextD < dp[nextI][k] {
					dp[nextI][k] = nextD
				}
			}
		}
	}

	all := 1<<n - 1
	ret := math.MaxFloat64

	for i := uint(0); i < n; i++ {
		if dp[all][i] == math.MaxFloat64 {
			continue
		}

		temp := dp[all][i] + dist[i][0]
		if temp < ret {
			ret = temp
		}
	}
	return ret
}
