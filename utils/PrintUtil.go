package utils

import (
	"fmt"
	"github.com/benjaminliu/leetcode-go-2/structures"
)

func PrintListNode(head *structures.ListNode) {
	for head != nil {
		fmt.Printf("%d ->", head.Val)
		head = head.Next
	}
}

func Println(num int) {
	fmt.Println(num)
}
