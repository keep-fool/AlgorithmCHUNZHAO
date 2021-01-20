package week1

import "testing"

func TestClimbingStairs(t *testing.T) {
	t.Log(climbStairs(5))
}

// 斐波那契
func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	v1, v2, res := 1, 1, 2
	for i := 2; i <= n; i++ {
		res = v1 + v2
		v1, v2 = v2, res
	}
	return res
}
