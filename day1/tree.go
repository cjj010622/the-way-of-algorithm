package day1

import "math"

var res int = math.MinInt32

// 二叉树上最大路径和
func oneSideMax(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//左侧节点最大贡献值
	left := int(math.Max(0, float64(oneSideMax(root.Left))))
	//右侧节点最大贡献值
	right := int(math.Max(0, float64(oneSideMax(root.Right))))
	//计算经过当前节点的路径和最大值
	//后序遍历
	res = int(math.Max(float64(res), float64(left+right+root.Val)))

	return int(math.Max(float64(left), float64(right))) + root.Val
}

// 根据前序遍历和中序遍历构建二叉树
func build(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	var index int

	//在中序遍历中获取根节点的位置
	for i := inStart; i <= inEnd; i++ {
		if root.Val == inorder[i] {
			index = i
			break
		}
	}

	//左子树的长度
	leftSize := index - inStart
	//前序遍历只取到左子树的长度
	root.Left = build(preorder, preStart+1, preStart+leftSize,
		//inStart-1是上面for循环i取到inEnd
		inorder, inStart, index-1)

	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd)

	return root
}

//获取二叉树中第k小的数

var rank int
var res2 int

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}

	traverse(root.Left, k)
	//中序遍历位置
	rank++
	if k == rank {
		res2 = root.Val
		return
	}

	traverse(root.Right, k)
}
