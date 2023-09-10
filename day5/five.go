package day5

//多数元素

// 方式一
func majorityElement(nums []int) int {
	m := make(map[int]int)
	l := len(nums) / 2
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > l {
			return nums[i]
		}
	}
	return 0
}

//方式二
//摩尔投票法
func majorityElement2(nums []int) int {
	r, c := 0, 0
	for _, v := range nums {
		if c == 0 {
			r = v
		}

		if r == v {
			c++
		} else {
			c--
		}
	}

	return r
}
