package week03

import "testing"

func TestSearchMatrix(t *testing.T) {

	var (
		matrix = [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target = 3
	)
	t.Log(searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	line := -1
	for i := 0; i < len(matrix)-1; i++ {
		if target >= matrix[i][0] && target < matrix[i+1][0] {
			line = i
			break
		}
	}
	if line == -1 && target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	if line == -1 {
		line = len(matrix) - 1
	}

	right := len(matrix[0]) - 1
	left := 0
	for left <= right {
		mid := left + (right-left)/2
		if matrix[line][mid] == target {
			return true
		}
		if matrix[line][mid] < target {
			left = mid + 1
		} else if matrix[line][mid] > target {
			right = mid - 1
		}
	}
	return false
}
