package day23

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。

如果可以，返回 true ；否则返回 false 。

magazine 中的每个字符只能在 ransomNote 中使用一次。
*/

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		count, ok := m[ransomNote[i]]
		if !ok {
			return false
		}

		count--
		m[ransomNote[i]] = count
		if count < 0 {
			return false
		}
	}
	return true
}
