package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	newHead, newTail := reverseList(head, cur)
	newTail.Next = reverseKGroup(cur, k)
	return newHead
}

func reverseList(from, to *ListNode) (*ListNode, *ListNode) {
	var pre, cur, nxt *ListNode
	pre = nil
	cur = from
	for cur != to {
		nxt = cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre, from
}
