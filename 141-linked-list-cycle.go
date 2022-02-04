package main

// https://leetcode-cn.com/problems/linked-list-cycle/

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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	for fast, slow := head.Next, head; fast != slow; {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
