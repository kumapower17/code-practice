package main

// https://leetcode-cn.com/problems/insertion-sort-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(head, node *ListNode) {
	var p = head
	for p.Next != nil && p.Next.Val < node.Val {
		p = p.Next
	}
	node.Next = p.Next
	p.Next = node
}

func insertionSortList(head *ListNode) *ListNode {
	var dummy ListNode
	for p := head; p != nil; {
		next := p.Next
		insert(&dummy, p)
		p = next
	}
	return dummy.Next
}
