package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	//有符号的最大最小值
	const INT_MAX = int(^uint(0) >> 1)

	minDiff := INT_MAX
	pre := -1
	inorder(root, &minDiff, &pre)
	return minDiff

}

func inorder(p *TreeNode, minDiff *int, pre *int) {
	if p.Left != nil {
		inorder(p.Left, minDiff, pre)
	}
	if *pre != -1 {
		*minDiff = int(math.Min((float64(*minDiff)), float64(p.Val-*pre)))
	}
	*pre = p.Val
	if p.Right != nil {
		inorder(p.Right, minDiff, pre)
	}
}
