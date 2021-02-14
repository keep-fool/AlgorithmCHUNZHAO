package week04

import "testing"

func TestMinimumTotal(t *testing.T) {

	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	t.Log(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	bottom := triangle[len(triangle)-1]
	dp := make([]int, len(bottom))
	for i := range dp {
		dp[i] = bottom[i]
	}
	for i := len(dp) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}
