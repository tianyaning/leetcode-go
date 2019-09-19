package main

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isUnivalTree1(root, root.Val)

}

func isUnivalTree1(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return isUnivalTree1(root.Left, val) && isUnivalTree1(root.Right, val)
}
