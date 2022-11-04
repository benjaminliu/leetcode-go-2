package main

import "github.com/benjaminliu/leetcode-go-2/structures"

type ListNode = structures.ListNode

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB

	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}

	return curA
}
