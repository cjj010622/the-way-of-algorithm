package day1

// 数组迭代
func arrTraverse(arr []int) {
	for i := 0; i < len(arr); i++ {
		//......
	}
}

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表迭代
func linkListTraverse(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		// ......
	}
}

// 递归迭代链表
func linkListTraverseRecursively(head *ListNode) {
	linkListTraverseRecursively(head.Next)
}

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 基本二叉树迭代
func treeTraverse(root *TreeNode) {
	if root != nil {
		treeTraverse(root.Left)
		treeTraverse(root.Right)
	}
}

// NTreeNode N叉树节点
type NTreeNode struct {
	Val      int
	Children []*NTreeNode
}

// 基本n叉树迭代
func nTreeTraverse(root *NTreeNode) {
	for _, child := range root.Children {
		nTreeTraverse(child)
	}
}
