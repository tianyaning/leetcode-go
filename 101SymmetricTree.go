package main

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}
	return isSymmetricChild(root.Left, root.Right)

}

func isSymmetricChild(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return (p.Val == q.Val) && isSymmetricChild(p.Left, q.Right) && isSymmetricChild(p.Right, q.Left)
}
