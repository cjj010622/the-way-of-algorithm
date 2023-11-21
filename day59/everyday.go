package day59

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 k 。你需要执行以下操作 恰好 k 次，最大化你的得分：

从 nums 中选择一个元素 m 。
将选中的元素 m 从数组中删除。
将新元素 m + 1 添加到数组中。
你的得分增加 m 。
请你返回执行以上操作恰好 k 次后的最大得分。
*/

func maximizeSum(nums []int, k int) int {
	x := slices.Max(nums)
	return k*x + k*(k-1)/2
}
