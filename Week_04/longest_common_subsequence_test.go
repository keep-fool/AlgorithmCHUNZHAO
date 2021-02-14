package week04

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	var text1, text2 = "abcde", "ace"
	t.Log(longestCommonSubsequence(text1, text2))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
			} else {
				dp[j] = func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(tmp, dp[j-1])
			}
			last = tmp
		}
	}
	return dp[len(text2)]
}

func longestCommonSubsequence1(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
			} else {
				dp[j] = func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(tmp, dp[j-1])
			}
			last = tmp
		}
	}
	return dp[len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
