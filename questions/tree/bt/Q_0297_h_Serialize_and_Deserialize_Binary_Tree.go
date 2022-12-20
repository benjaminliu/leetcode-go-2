package main

import (
	"github.com/benjaminliu/leetcode-go-2/structures"
	"strconv"
	"strings"
)

type TreeNode = structures.TreeNode

func main() {
	ser := Constructor()
	deser := Constructor()

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	data := ser.serialize(root)

	ans := deser.deserialize(data)

	println(ans)
}

type Codec struct {
	SEP       string
	NULL      string
	decodeIdx int
	sb        strings.Builder
}

func Constructor() Codec {
	return Codec{SEP: ",", NULL: "#"}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.sb = strings.Builder{}
	this.traverse(root)
	return this.sb.String()
}

func (this *Codec) traverse(root *TreeNode) {
	if root == nil {
		this.sb.WriteString(this.NULL)
		this.sb.WriteString(this.SEP)
		return
	}

	this.sb.WriteString(strconv.Itoa(root.Val))
	this.sb.WriteString(this.SEP)

	this.traverse(root.Left)
	this.traverse(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.decodeIdx = 0

	nodes := strings.Split(data, this.SEP)

	return this.buildTree(nodes)
}

func (this *Codec) buildTree(nodes []string) *TreeNode {
	if this.decodeIdx >= len(nodes) {
		return nil
	}
	val := nodes[this.decodeIdx]
	this.decodeIdx++
	if this.NULL == val {
		return nil
	}

	valInt, _ := strconv.Atoi(val)
	root := &TreeNode{Val: valInt}

	root.Left = this.buildTree(nodes)
	root.Right = this.buildTree(nodes)
	return root
}
