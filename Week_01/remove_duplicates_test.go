package week01

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var nums = []int{1, 1, 2}
	t.Log(removeDuplicates(nums))
}

// 快慢指针
func removeDuplicates(nums []int) int {
	lens := len(nums)
	var i, j = 0, 1
	for j < lens {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)
	return i + 1
}
