package day20

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	rk, ans := -1, 0
	n := len(s)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
