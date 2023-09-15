package day10

// 跳跃游戏二
func jump(nums []int) int {
	l := len(nums)
	maxPosition := 0
	end := 0
	step := 0
	for i := 0; i < l-1; i++ {
		maxPosition = max(maxPosition, nums[i]+i)
		if i == end {
			end = maxPosition
			step++
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
