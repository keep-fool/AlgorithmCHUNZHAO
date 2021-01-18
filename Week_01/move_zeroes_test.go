package week1

import "testing"

func TestMoveZeroes(t *testing.T) {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes1(nums)
}

// moveZeroes1 解法一 双指针
func moveZeroes1(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			i++
		}
	}
}
