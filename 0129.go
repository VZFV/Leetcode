package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
1. Base Case    global_sum += sum
2. 利用父问题传下来的值做一些计算   concat digits: num += num * 10 + node.val
3. 若有必要，做一些额外操作
4. 把值传下去给子问题继续递归    dfs(child_node, num)
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	dfs(root, 0, &max)
	return max
}

func dfs(root *TreeNode, num int, max *int) {
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*max += num
	}
	if root.Left != nil {
		dfs(root.Left, num, max)
	}
	if root.Right != nil {
		dfs(root.Right, num, max)
	}
}

