package main

// https://leetcode-cn.com/problems/rotate-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

func iter(head, node *ListNode, count, k *int) *ListNode {
	if node == nil {
		return head
	}
	*count++

	if node.Next == nil {
		*k %= (*count)
	}

	h := iter(head, node.Next, count, k)

	if *k == 0 {
		node.Next = nil
	} else if *k > 0 {
		node.Next = h
		h = node
	}
	*k--
	return h
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	k1, count := k, 0
	return iter(head, head, &count, &k1)
}
