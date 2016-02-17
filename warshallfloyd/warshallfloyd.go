package warshallfloyd

import "fmt"

// Search is Warshall Floyd algorithm
func Search(dp [][]int) ([][]int, error) {
	for k := 0; k < len(dp); k++ {
		for i := 0; i < len(dp); i++ {
			for j := 0; j < len(dp); j++ {
				cost := dp[i][j]
				other := dp[i][k] + dp[k][j]
				if other < cost {
					cost = other
				}

				dp[i][j] = cost
			}
		}
	}

	for i := range dp {
		if dp[i][i] < 0 {
			return dp, fmt.Errorf("negative cycle")
		}
	}

	return dp, nil
}

// OptRoute is optimal routing
func OptRoute(dp [][]int) ([][]int, error) {
	for k := 0; k < len(dp); k++ {
		for i := 0; i < len(dp); i++ {
			for j := 0; j < len(dp); j++ {
				cost := dp[i][j]
				other := dp[i][k]
				if dp[k][j] < other {
					other = dp[k][j]
				}

				if other > cost {
					cost = other
				}

				dp[i][j] = cost
			}
		}
	}

	for i := range dp {
		if dp[i][i] < 0 {
			return dp, fmt.Errorf("negative cycle")
		}
	}

	return dp, nil
}
