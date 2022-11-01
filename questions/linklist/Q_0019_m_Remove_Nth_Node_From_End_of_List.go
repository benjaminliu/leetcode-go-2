package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	for n >= 0 {
		fast = fast.Next
		n--
	}
	slow := dummy
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
