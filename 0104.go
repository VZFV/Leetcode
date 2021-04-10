package main
/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	max_depth := max(left, right) + 1
	return max_depth
}
