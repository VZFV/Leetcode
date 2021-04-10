package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
Bottom Up DFS 解题套路：
1. Base case (如root == nil, return 0)
2. 向子问题要答案 (max sum root to leaf path)
3. 利用子问题的答案构建当前问题(当前递归层)的答案 ans = max(left,right,0)+node.val
4. 如果有必要，做一些额外操作(比如及时更新global最大值) global_max = max(left,0) + max(right,0)+node.val
5. 返回答案(给父问题)  return ans
 */
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32 // set the minimum number to max
	getPathSum(root, &max)
	return max
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	currMax := max(max(left+root.Val, right+root.Val), root.Val) // currMax = max(left,right,0) + root.val
	*maxSum = max(*maxSum, max(currMax, max(left,0)+max(right,0)+root.Val)) // update global max
	return currMax
}