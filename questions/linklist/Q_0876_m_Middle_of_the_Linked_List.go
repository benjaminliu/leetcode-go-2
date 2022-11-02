package main

import "github.com/benjaminliu/leetcode-go-2/structures"

func main() {

}

type ListNode = structures.ListNode

func middleNode(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	fast := dummy
	slow := dummy

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow.Next
}
