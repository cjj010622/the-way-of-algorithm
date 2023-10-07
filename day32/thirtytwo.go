package day32

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
