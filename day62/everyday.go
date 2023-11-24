package day62

import "sort"

/*
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target ，请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目。
*/

func countPairs(nums []int, target int) int {
	var ans int
	sort.Ints(nums)
	for j,x:=range nums{
		i:=sort.SearchInts(nums[:j],target-x)
		ans+=i
	}
	return ans
}