package day13

import "sort"

//最长公共前缀

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	d, c := strs[0], strs[len(strs)-1]
	i := 0
	for ; i < len(d); i++ {
		if d[i] != c[i] {
			break
		}
	}
	return d[:i]
}
