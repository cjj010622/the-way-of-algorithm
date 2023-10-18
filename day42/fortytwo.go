package day42

import "math"

/*
给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt32, -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)
		if pre != -1 && node.Val-pre < res {
			res = node.Val - pre
		}

		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)

	return res
}
