package day58

/*
给定一个长度为 n 的环形整数数组 nums ，返回 nums 的非空 子数组 的最大可能和 。

环形数组 意味着数组的末端将会与开头相连呈环状。形式上， nums[i] 的下一个元素是 nums[(i + 1) % n] ， nums[i] 的前一个元素是 nums[(i - 1 + n) % n] 。

子数组 最多只能包含固定缓冲区 nums 中的每个元素一次。形式上，对于子数组 nums[i], nums[i + 1], ..., nums[j] ，不存在 i <= k1, k2 <= j 其中 k1 % n == k2 % n 。
*/

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	leftMax := make([]int, n)

	leftMax[0] = nums[0]
	leftSum, pre, res := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		pre = max(pre+nums[i], nums[i])
		res = max(res, pre)
		leftSum += nums[i]
		leftMax[i] = max(leftMax[i-1], leftSum)
	}

	rightSum := 0
	for i := n - 1; i > 0; i-- {
		rightSum += nums[i]
		res = max(res, rightSum+leftMax[i-1])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
