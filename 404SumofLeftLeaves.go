package main

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	dfSearch(root, &sum, false)
	return sum
}

//到leaf的时候判断是不是left leaf。用一个boolean function来记录搜索的是左枝还是右枝。
func dfSearch(root *TreeNode, sum *int, left bool) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if left {
			*sum += root.Val
		}
	}
	dfSearch(root.Left, sum, true)
	dfSearch(root.Right, sum, false)
}
