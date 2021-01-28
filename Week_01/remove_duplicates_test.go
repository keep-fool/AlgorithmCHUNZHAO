package week01

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// {
	// 	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// 	// var nums = []int{1, 1, 2}
	// 	t.Log(removeDuplicates(nums))
	// 	t.Log(nums)
	// }
	{
		var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		// var nums = []int{1, 1, 2}
		t.Log(removeDuplicates1(nums))
		t.Log(nums)
	}
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
			j++
		}
	}
	return i + 1
}

func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return i + 1
}

func removeDuplicates2(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
