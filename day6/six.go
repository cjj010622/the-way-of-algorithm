package day6

//轮转数组

// 方式一
func rotate(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i, v := range nums {
		tmp[(i+k)%len(nums)] = v
	}
	copy(nums, tmp)
}

// 方式二
func rotate2(nums []int, k int) {
	tmp := make([]int, len(nums))
	for i, v := range nums {
		tmp[i] = v
	}
	for i := 0; i < len(nums); i++ {
		nums[(i+k)%len(nums)] = tmp[i]
	}
}

// 方式三
func reverse(arr []int) {
	i, n := 0, len(arr)
	for ; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func rotate3(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
