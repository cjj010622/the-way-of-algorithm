package day1

// 全排列算法
func backtrack(nums []int, track []int, res *[][]int) {
	//判断是否达到决策树底层
	if len(nums) == len(track) {
		*res = append(*res, append([]int{}, track...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}

		track = append(track, nums[i])
		backtrack(nums, track, res)
		track = track[:len(track)-1]
	}
}

func contains(track []int, val int) bool {
	for _, v := range track {
		if v == val {
			return true
		}
	}
	return false
}
