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
	t.Log(preorderTraversal(root))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 中序遍历
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

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preorder(root, &res)
	return res
}

func preorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorder(root.Left, res)
	preorder(root.Right, res)
}
