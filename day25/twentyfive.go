package day25

import "strings"

/*
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", s = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", s = "dog cat cat fish"
输出: false
*/

func wordPattern(pattern string, s string) bool {
	p2s := make(map[byte]string)
	s2p := make(map[string]byte)
	ps := strings.Split(s, " ")
	if len(ps) != len(pattern) {
		return false
	}
	for i := 0; i < len(pattern); i++ {
		a := ps[i]
		b := pattern[i]
		if p2s[b] != "" && p2s[b] != a || s2p[a] != 0 && s2p[a] != b {
			return false
		}

		p2s[b] = a
		s2p[a] = b
	}
	return true
}
