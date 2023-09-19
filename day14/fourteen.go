package day14

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
*/

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	la := make([]int, n)
	ra := make([]int, n)

	la[0] = 1
	for i := 1; i < n; i++ {
		la[i] = la[i-1] * nums[i-1]
	}

	ra[n-1] = 1
	for j := n - 2; j >= 0; j-- {
		ra[j] = ra[j+1] * nums[j+1]
	}

	for i := 0; i < n; i++ {
		res[i] = la[i] * ra[i]
	}

	return res
}
