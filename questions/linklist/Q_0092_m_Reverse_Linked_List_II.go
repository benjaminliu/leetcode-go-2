package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		newHead, _ := reverseFirstN(head, right)
		return newHead
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseFirstN(head *ListNode, n int) (*ListNode, *ListNode) {
	if n == 1 {
		return head, head.Next
	}

	next := head.Next

	newHead, successor := reverseFirstN(head.Next, n-1)

	next.Next = head
	head.Next = successor

	return newHead, successor
}
