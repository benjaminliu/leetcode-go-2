package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return partition(lists, 0, len(lists)-1)
}

func partition(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}

	mid := (start + end) / 2
	left := partition(lists, start, mid)
	right := partition(lists, mid+1, end)

	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	dummy := &ListNode{}
	cur := dummy

	for left != nil && right != nil {
		if left.Val <= right.Val {
			cur.Next = left
			cur = cur.Next
			left = left.Next
		} else {
			cur.Next = right
			cur = cur.Next
			right = right.Next
		}
	}
	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return merge(lists[0], lists[1])
	}

	mid := len(lists) / 2
	left := mergeKLists1(lists[:mid])
	right := mergeKLists1(lists[mid:])
	return merge(left, right)
}
