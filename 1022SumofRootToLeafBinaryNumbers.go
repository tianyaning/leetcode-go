package main

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	return f(root, sum)
}

func f(t *TreeNode, sum int) int {
	if t.Left == nil && t.Right == nil {
		//叶子结点，直接计算结果
		return sum*2 + t.Val
	} else {
		if t.Left != nil && t.Right != nil {
			return f(t.Left, sum*2+t.Val) + f(t.Right, sum*2+t.Val)
		} else if t.Left != nil {
			return f(t.Left, sum*2+t.Val)
		} else {
			return f(t.Right, sum*2+t.Val)
		}
	}
}
