package week01

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	var list = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	var list1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	res := mergeTwoLists(list, list1)
	t.Logf("%+v", *res)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
		fmt.Println(prehead)
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
