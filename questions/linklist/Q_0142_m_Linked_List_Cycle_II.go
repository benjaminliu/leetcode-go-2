package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast, slow := dummy, dummy

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = dummy
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}

			return fast
		}
	}
	return nil
}
