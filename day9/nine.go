package day9

//跳跃游戏

func canJump(nums []int) bool {
	n := len(nums)
	r := 0
	for i := 0; i < n-1; i++ {
		r = max(r, nums[i]+i)

		if r < i {
			return false
		}
	}

	return r >= n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
