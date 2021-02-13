package week04

import (
	"testing"
)

func TestMinimumPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	t.Log(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}
	// 初始化DP
	dp := make([][]int, len(grid))
	for k, v := range grid {
		dp[k] = make([]int, len(v))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch {
			case i == 0 && j != 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case i != 0 && j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
				// 最优
			case i != 0 && j != 0:
				dp[i][j] = func(a, b int) int {
					if a > b {
						return b
					}
					return a
				}(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func minPathSum1(grid [][]int) int {
	if grid == nil {
		return 0
	}
	dp := make([][]int, len(grid))
	for k, v := range grid {
		dp[k] = make([]int, len(v)) //初始化dp数组
	}
	dp[0][0] = grid[0][0] //动态规划边界
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j] //动态规划递推关系式
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j] //动态规划递推关系式
			} else if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j] ////动态规划递推关系式
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
