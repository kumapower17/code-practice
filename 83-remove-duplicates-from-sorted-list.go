package main

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p, q := head, head.Next
	for q != nil {
		if q.Val == p.Val {
			p.Next = q.Next
			q.Next = nil
			q = p.Next
		} else {
			p = q
			q = q.Next
		}
	}
	return head
}
