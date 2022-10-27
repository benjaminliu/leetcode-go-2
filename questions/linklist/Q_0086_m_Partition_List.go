package main

import (
	"github.com/benjaminliu/leetcode-go-2/structures"
	"github.com/benjaminliu/leetcode-go-2/utils"
)

type ListNode = structures.ListNode

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 3}

	res := partition(head, 3)
	utils.PrintListNode(res)
}

func partition(head *ListNode, x int) *ListNode {
	dummySmall := &ListNode{}
	dummyBig := &ListNode{}

	curSmall := dummySmall
	curBig := dummyBig

	cur := head

	for cur != nil {
		tmp := cur
		cur = cur.Next
		tmp.Next = nil

		if tmp.Val < x {
			curSmall.Next = tmp
			curSmall = curSmall.Next
		} else {
			curBig.Next = tmp
			curBig = curBig.Next
		}
	}

	curSmall.Next = dummyBig.Next
	return dummySmall.Next
}
