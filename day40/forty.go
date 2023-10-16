package day40

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方式一
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parants := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			parants[root.Left.Val] = root
			dfs(root.Left)
		}

		if root.Right != nil {
			parants[root.Right.Val] = root
			dfs(root.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p.Val] = true
		p = parants[p.Val]
	}

	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parants[q.Val]
	}

	return nil
}

// 方式二
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}
