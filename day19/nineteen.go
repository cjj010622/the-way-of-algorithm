package day19

import "math"

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/

// 暴力
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum >= target {
			return 1
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				res = min(res, j+1-i)
				break
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n {
		sum += nums[end]
		for sum >= target {
			res = min(res, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}
