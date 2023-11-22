package day60

/*
给你一个下标从 0 开始的整数数组 nums 和一个整数 threshold 。

请你从 nums 的子数组中找出以下标 l 开头、下标 r 结尾 (0 <= l <= r < nums.length) 且满足以下条件的 最长子数组 ：

nums[l] % 2 == 0
对于范围 [l, r - 1] 内的所有下标 i ，nums[i] % 2 != nums[i + 1] % 2
对于范围 [l, r] 内的所有下标 i ，nums[i] <= threshold
以整数形式返回满足题目要求的最长子数组的长度。

注意：子数组 是数组中的一个连续非空元素序列。
*/

func longestAlternatingSubarray(nums []int, threshold int) int {
	var ans int
	for l,n:=0,len(nums);l<n;{
		if nums[l]%2==0&&nums[l]<=threshold{
			r:=l+1
			for r<n&&nums[r]%2!=nums[r-1]%2&&nums[r]<=threshold{
				r++
			}
			ans=max(ans,r-l)
			l=r
		}else{
			l++
		}
	}
	return ans
}