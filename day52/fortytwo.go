package day52

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	used := make([]bool, len(nums))
	var tmp []int

	var res [][]int
	var backTrack func(index int)
	backTrack = func(index int) {
		if len(tmp) == len(nums) {
			res = append(res, append([]int{}, tmp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !used[i] {
				used[i] = true
				tmp[index] = nums[i]
				backTrack(index + 1)
				used[i] = false
			}
		}

	}
	backTrack(0)
	return res
}
