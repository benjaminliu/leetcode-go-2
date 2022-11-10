package main

func main() {

}

var left *ListNode

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	left = head
	return traverse(head.Next)
}

func traverse(right *ListNode) bool {
	if left == right {
		return true
	}

	if right.Next != nil {
		res := traverse(right.Next)
		if res == false {
			return false
		}
	}

	if left.Val != right.Val {
		return false
	}

	left = left.Next
	return true
}
