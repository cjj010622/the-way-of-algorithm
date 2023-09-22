package day17

import "sort"

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	l := len(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left := i + 1
		right := l - 1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
