package week03

import "testing"

func TestMajorityElement(t *testing.T) {

	var nums = []int{2, 2, 1, 1, 1, 2, 2}

	t.Log(majorityElement(nums))
}

// 暴力
func majorityElement(nums []int) int {
	maps := make(map[int]int)
	for _, v := range nums {
		maps[v]++
	}
	var out, max int
	for k, v := range maps {
		if v > max {
			max = v
			out = k
		}
	}
	return out
}

// 投票法
func majorityElement1(nums []int) int {
	major := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}

	return major
}
