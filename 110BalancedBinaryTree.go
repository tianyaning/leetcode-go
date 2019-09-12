package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		h1 := dfs(root.Left)
		h2 := dfs(root.Right)
		if h1-h2 > 1 || h1-h2 < -1 {
			return false
		} else {
			return isBalanced(root.Left) && isBalanced(root.Right)
		}
	}
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		h1 := dfs(root.Left) + 1
		h2 := dfs(root.Right) + 1
		if h1 >= h2 {
			return h1
		} else {
			return h2
		}
	}
}
