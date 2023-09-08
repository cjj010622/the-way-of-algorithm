package day1

//最深叶节点的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, r := dfs(root)
	return r
}

func dfs(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	d1, lf := dfs(root.Left)
	d2, rg := dfs(root.Right)

	if d1 > d2 {
		return d1 + 1, lf
	}

	if d2 > d1 {
		return d2 + 1, rg
	}

	return d1 + 1, root
}
