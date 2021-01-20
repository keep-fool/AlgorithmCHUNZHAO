package week1

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	var height = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	{
		res := maxArea1(height)
		t.Log(res)
	}
	{
		res := maxArea1(height)
		t.Log(res)
	}
}

// 解法一 暴力 双重遍历
func maxArea1(height []int) int {
	var max = 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			min := func(a, b int) int {
				if a > b {
					return b
				}
				return a
			}(height[i], height[j])
			if max < min*(j-i) {
				max = min * (j - i)
			}
		}
	}
	return max
}

// 解法二 双指针 夹逼
func maxArea2(height []int) int {
	o := 0
	i, j := 0, len(height)-1
	for i != j {
		hi, hj := height[i], height[j]
		s := (j - i) * func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}(height[i], height[j])

		if s > o {
			o = s
		}
		if hi > hj {
			j--
		} else {
			i++
		}
	}
	return o
}
