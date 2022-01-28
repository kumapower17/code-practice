package main

// https://leetcode-cn.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := ListNode{
		Next: head,
	}
	var leftNodeParent = &dummy
	for i := 1; i < left; i++ {
		leftNodeParent = leftNodeParent.Next
	}
	leftNode := leftNodeParent.Next

	for i := left; i < right; i++ {
		node := leftNode.Next
		leftNode.Next = node.Next
		node.Next = leftNodeParent.Next
		leftNodeParent.Next = node
	}
	return dummy.Next
}
