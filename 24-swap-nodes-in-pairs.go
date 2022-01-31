package main

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/submissions/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// two or more nodes
	a, b := head, head.Next
	remaining := swapPairs(b.Next)
	a.Next = remaining
	b.Next = a
	return b
}
