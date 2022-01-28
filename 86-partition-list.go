package main

// https://leetcode-cn.com/problems/partition-list/

type ListNode struct {
	Val  int
	Next *ListNode
}

type dList struct {
	head *ListNode
	tail *ListNode
}

func (dl *dList) append(node *ListNode) {
	if dl.head == nil {
		dl.head = node
	}
	if dl.tail != nil {
		dl.tail.Next = node
	}
	dl.tail = node
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var smaller, notSmaller dList

	for node := head; node != nil; {
		next := node.Next
		node.Next = nil
		if node.Val < x {
			smaller.append(node)
		} else {
			notSmaller.append(node)
		}
		node = next
	}

	// smaller is empty
	if smaller.head == nil {
		return notSmaller.head
	}

	// join two list
	smaller.tail.Next = notSmaller.head
	return smaller.head
}
