package day51

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。
*/

func combine(n int, k int) [][]int {
	var na []int
	for i := 1; i <= n; i++ {
		na = append(na, i)
	}
	var tmp []int

	var res [][]int
	var backTrack func(index int)
	backTrack = func(index int) {
		if len(tmp) == k {
			res = append(res, append([]int{}, tmp...))
			return
		}
		for i := index; i < n; i++ {
			tmp = append(tmp, na[i])
			backTrack(i + 1)
			tmp = tmp[:len(tmp)-1]
		}
	}

	backTrack(0)
	return res
}
