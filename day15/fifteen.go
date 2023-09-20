package day15

import "strings"

/*
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。

示例 1：
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
示例 2：

示例 2：
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
*/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removeNonAlphanumeric(s)

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func removeNonAlphanumeric(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= '0' || c <= '9') {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
