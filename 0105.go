package main

// the first element of preorder array must be root. Then we can find the index of the same element in
// inoder array, any element in the left is left tree. right is right tree
func buildTree(preorder []int, inorder []int) *TreeNode {
	// create a map that store the value and index of inorder array
	inPos := make(map[int]int)

	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return helper(preorder, inorder, 0, 0, len(inorder)-1, inPos)
}

func helper(preorder []int, inorder []int, preStart int, inStart int, inEnd int, inPos map[int]int) *TreeNode {
	if preStart >= len(preorder) || inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		Val: preorder[preStart],
	} // find the root of preorder array
	rootIndex := inPos[preorder[preStart]]
	root.Left = helper(preorder, inorder, preStart+1, inStart, rootIndex-1, inPos)
	root.Right = helper(preorder, inorder, rootIndex-inStart+1+preStart, rootIndex+1, inEnd, inPos)
	return root
}

/** this method saves more memory, don't need a map to store index and value
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(inorder) != len(preorder) || len(preorder) == 0 || len(preorder) == 0 {
        return nil
    }
    rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}
    i := 0
    for i = 0; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
 */