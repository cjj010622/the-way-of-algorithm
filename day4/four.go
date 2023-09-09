package day4

////去除数组中重复两次以上数

// 方式一 使用map
func removeDuplicates(nums []int) int {
	i := 0
	m := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		m[nums[j]]++
		if c := m[nums[j]]; c < 3 {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 方式二 原地
func removeDuplicates2(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}

	i := 2
	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
