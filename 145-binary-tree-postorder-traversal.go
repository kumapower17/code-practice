package main

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iter(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}
	values = iter(root.Left, values)
	values = iter(root.Right, values)
	return append(values, root.Val)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	return iter(root, nil)
}
