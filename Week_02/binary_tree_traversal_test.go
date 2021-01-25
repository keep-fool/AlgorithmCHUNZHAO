package week02

import (
	"fmt"
	"testing"
)

func TestBinaryTreeTraversal(t *testing.T) {
	var root = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	//   4
	// 2   6
	//1 3 5 7
	t.Log(inorderTraversal(root))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	inorder(root, &res)
	return res
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}
