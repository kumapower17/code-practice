package main

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iter(root *TreeNode, values []int) []int {
	if root == nil {
		return values
	}
	values = append(values, root.Val)
	values = iter(root.Left, values)
	return iter(root.Right, values)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var vals []int
	return iter(root, vals)
}
