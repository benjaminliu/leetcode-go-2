package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	connectTwo(root.Left, root.Right)

	return root
}

func connectTwo(left, right *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right

	connectTwo(left.Left, left.Right)
	connectTwo(left.Right, right.Left)
	connectTwo(right.Left, right.Right)
}
