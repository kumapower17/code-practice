package main

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
	values = append(values, root.Val)
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
func inorderTraversal(root *TreeNode) []int {
	return iter(root, nil)
}
