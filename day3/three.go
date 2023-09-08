package day3

//去除数组中重复数

// 方式一：
// 如果不是递增，通过map去重
func removeDuplicates(nums []int) int {
	i := 0
	m := make(map[int]bool)
	for j := 0; j < len(nums); j++ {
		if m[nums[j]] {
			continue
		}
		nums[i] = nums[j]
		i++
		m[nums[j]] = true
	}
	return i
}

// 方式二
// 这是递增数组 与前一个相比
func removeDuplicates2(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
