package day61

/*
给你一个仅由 0 和 1 组成的二进制字符串 s 。

如果子字符串中 所有的 0 都在 1 之前 且其中 0 的数量等于 1 的数量，则认为 s 的这个子字符串是平衡子字符串。请注意，空子字符串也视作平衡子字符串。

返回  s 中最长的平衡子字符串长度。

子字符串是字符串中的一个连续字符序列。
*/

func findTheLongestBalancedSubstring(s string) (ans int) {
	zero, one := 0, 0
	for _, c := range s {
		if c == '0' {
			if one > 0 {
				zero, one = 0, 0
			}
			zero++
		} else {
			one++
			ans = max(ans, 2*min(zero, one))
		}
	}
	return
}
