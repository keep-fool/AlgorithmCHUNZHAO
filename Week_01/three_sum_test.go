package week01

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	t.Log(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	nums = sort.IntSlice(nums)
	var out [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for m := j + 1; m < len(nums); m++ {
				if nums[i]+nums[j]+nums[m] == 0 {
					// func(n1, n2, n3 int) []int {
					// }(nums[i], nums[j], nums[m])
					out = append(out, []int{nums[i], nums[j], nums[m]})
				}
			}
		}
	}

	return out
}
