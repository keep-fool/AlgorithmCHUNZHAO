package week02

import (
	"sort"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	var (
		arr = []int{3, 2, 1}
		k   = 2
	)
	t.Log(getLeastNumbers(arr, k))
}

// 直接排序 返回前K个
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
