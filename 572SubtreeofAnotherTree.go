package main

//如果 s 和 t 一摸一样，那么 t 也是 s 的子树。
// 如果s和t不一样， 那么我们要依次从 s 的 left 当作一个新的树，和 t 比较，如果不是；从 s 的 right 当作一个新的树，和 t 比较。直到把 s 树遍历结束；
// 其中如果遇到 s 等于 t 的情况，立即返回true即可。
// 那么我们要另外设一个function，来判断一下是否两个二叉树是相同的。
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if isSame(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
