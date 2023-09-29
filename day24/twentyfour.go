package day24

/*
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
*/

// 方法一
func isIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	t2s := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		x := s[i]
		y := t[i]
		if s2t[x] != 0 && s2t[x] != y || t2s[y] != 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}

// 方法二
func isIsomorphic2(s string, t string) bool {
	sa := make([]int, 256)
	ta := make([]int, 256)
	for i := 0; i < len(s); i++ {
		a := s[i]
		b := t[i]
		if sa[a] != ta[b] {
			return false
		}

		sa[a] = i + 1
		ta[b] = i + 1
	}
	return true
}
