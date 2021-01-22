package week01

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	t.Log(threeSum3(nums))
}

// 暴力  理解题意
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var out [][]int
	var re = make(map[string]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return out
		}
		if i > 0 && nums[i] == nums[i-1] {
			return out
		}
		for j := i + 1; j < len(nums); j++ {
			for m := j + 1; m < len(nums); m++ {
				if nums[i]+nums[j]+nums[m] == 0 {
					key := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[m])
					if _, ok := re[key]; !ok {
						re[key] = 0
						out = append(out, []int{nums[i], nums[j], nums[m]})
					}
				}
			}
		}
	}
	return out
}

// 双指针
func threeSum3(nums []int) [][]int {
	sort.Ints(nums)
	var (
		out  [][]int
		lens = len(nums)
	)
	if lens < 3 {
		return out
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return out
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针
		var j, m = i + 1, lens - 1
		for j < m {
			sum := nums[i] + nums[j] + nums[m]
			switch {
			case sum == 0:
				out = append(out, []int{nums[i], nums[j], nums[m]})
				for j < m && nums[j] == nums[j+1] {
					j += 1
				}
				for j < m && nums[m] == nums[m-1] {
					m -= 1
				}
				j, m = j+1, m-1
			case sum > 0:
				m -= 1
			case sum < 0:
				j += 1
			}
		}
	}
	return out
}

//
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	// 开始循环第一个固定值
	for index, value := range nums {
		// 如果固定位的值已经大于0，因为已经排好序了，后面的两个指针对应的值也肯定大于0，则和不可能为0，所以返回
		if nums[index] > 0 {
			return res
		}
		// 排除值重复的固定位
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		// 指针初始位置，固定位右边第一个和数组最后一个
		l := index + 1
		r := length - 1
		// 开始移动两个指针
		for l < r {
			// 判断三个数字之和的三种情况
			sum := value + nums[l] + nums[r]
			switch {
			case sum == 0:
				// 将结果加入二元组
				res = append(res, []int{nums[index], nums[l], nums[r]})
				// 去重，如果l < r且下一个数字一样，则继续挪动
				for l < r && nums[l] == nums[l+1] {
					l += 1
				}
				// 同理
				for l < r && nums[r] == nums[r-1] {
					r -= 1
				}
				l += 1
				r -= 1
			case sum > 0:
				// 如果和大于 0，那就说明 right 的值太大，需要左移
				r -= 1
				// 如果和小于 0，那就说明 left 的值太小，需要右移
			case sum < 0:
				l += 1
			}
		}
	}
	return res
}
