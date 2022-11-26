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
	fmt.Println()
}

func Println(num int) {
	fmt.Println(num)
}

func PrintList(list []int) {
	for _, val := range list {
		fmt.Printf("%d ->", val)
	}
	fmt.Println()
}
