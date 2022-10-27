package main

import (
	"github.com/benjaminliu/leetcode-go-2/structures"
)

func main() {

}

func mergeTwoLists(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	dummy := &structures.ListNode{}
	cur := dummy

	cur1 := list1
	cur2 := list2

	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	if cur1 != nil {
		cur.Next = cur1
	}

	if cur2 != nil {
		cur.Next = cur2
	}

	return dummy.Next
}
