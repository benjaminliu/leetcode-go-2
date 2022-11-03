package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
