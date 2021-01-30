package week02

import (
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
	t.Log(inorderTraversalBinary(root))
	t.Log(preorderTraversalBinary(root))
	t.Log(postorderTraversalBinary(root))
	t.Log(lowestCommonAncestor(root, root.Left.Right, root.Right))

}

// result

/*
$ go test -v Week_02/binary_tree_traversal_test.go
=== RUN   TestBinaryTreeTraversal
    binary_tree_traversal_test.go:32: [1 2 3 4 5 6 7]
    binary_tree_traversal_test.go:33: [4 2 1 3 6 5 7]
    binary_tree_traversal_test.go:34: [1 3 2 5 7 6 4]
--- PASS: TestBinaryTreeTraversal (0.00s)
PASS
ok      command-line-arguments  0.156s
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
// 中序遍历
func inorderTraversalBinary(root *TreeNode) []int {
	res := []int{}
	inorderBinary(root, &res)
	return res
}

func inorderBinary(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderBinary(root.Left, res)
	*res = append(*res, root.Val)
	inorderBinary(root.Right, res)
}

// 前序遍历
func preorderTraversalBinary(root *TreeNode) []int {
	res := []int{}
	preorderBinary(root, &res)
	return res
}

func preorderBinary(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderBinary(root.Left, res)
	preorderBinary(root.Right, res)
}

// 后序遍历
func postorderTraversalBinary(root *TreeNode) []int {
	res := []int{}
	postorderBinary(root, &res)
	return res
}

func postorderBinary(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorderBinary(root.Left, res)
	postorderBinary(root.Right, res)
	*res = append(*res, root.Val)
}

// 层寻遍历
func levelOrderBinary(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
