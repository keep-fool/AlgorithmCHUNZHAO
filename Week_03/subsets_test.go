package week03

import "testing"

func TestSubsets(t *testing.T) {
	var nums = []int{1, 2, 3}
	t.Log(subsets(nums))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		// 终止条件
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		// 保存结果
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})
	return res
}
