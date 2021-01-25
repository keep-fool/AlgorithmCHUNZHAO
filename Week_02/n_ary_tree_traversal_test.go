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
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

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
