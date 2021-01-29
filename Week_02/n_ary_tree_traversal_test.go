package week02

import "testing"

func TestPostorder(t *testing.T) {
	var root = &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	t.Log(postorder(root))
	t.Log(preorder(root))
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

// 中序遍历 递归
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	results := []int{}
	for _, v := range root.Children {
		results = append(results, postorder(v)...)
	}
	results = append(results, root.Val)
	return results
}

// 前序遍历  递归
func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var out []int
	out = append(out, root.Val)
	for _, node := range root.Children {
		out = append(out, preorder(node)...)
	}
	return out
}

func preorder1(root *Node) []int {
	var out []int
	var stack = []*Node{root}
	for len(stack) > 0 {
		for root != nil {
			out = append(out, root.Val)
			if len(root.Children) == 0 {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i])
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return out
}
