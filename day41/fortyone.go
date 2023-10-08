package day41

/*
给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type data struct {
	sum, count int
}

// dfs
func averageOfLevels(root *TreeNode) []float64 {
	var levelData []data
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level < len(levelData) {
			levelData[level].count++
			levelData[level].sum += root.Val
		} else {
			levelData = append(levelData, data{
				sum:   root.Val,
				count: 1,
			})
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)

	res := make([]float64, len(levelData))
	for i := 0; i < len(levelData); i++ {
		res[i] = float64(levelData[i].sum) / float64(levelData[i].count)
	}
	return res
}

// bfs
func averageOfLevels2(root *TreeNode) []float64 {
	var res []float64

	nextLevel := []*TreeNode{root}
	for len(nextLevel) > 0 {
		sum := 0
		curLevel := nextLevel
		nextLevel = nil
		for _, v := range curLevel {
			sum += v.Val
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}

		}
		res = append(res, float64(sum)/float64(len(curLevel)))
	}

	return res
}
