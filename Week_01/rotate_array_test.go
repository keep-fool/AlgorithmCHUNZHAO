package week01

import "testing"

func TestRotate(t *testing.T) {
	var (
		nums = []int{1, 2, 3, 4, 5, 6, 7}
		k    = 3
	)
	t.Log(rotate(nums, k))
}

func rotate(nums []int, k int) []int {
	res := []int{}
	res = append(nums[k:], nums[:k]...)
	nums = []int{}
	nums = res
	return nums
}

func rotate1(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
