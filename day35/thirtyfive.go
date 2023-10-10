package day35

/*
给定两个整数数组 inorder 和 postorder ，其中 inorder 是二叉树的中序遍历， postorder 是同一棵树的后序遍历，请你构造并返回这颗 二叉树 。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	root.Left = buildTree(inorder[:i], postorder[:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])

	return root
}
