package day50

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

*/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var combinations []string

	var tmp string
	var backtrack func(index int)
	backtrack = func(index int) {
		if len(tmp) == len(digits) {
			combinations = append(combinations, tmp)
		}

		for i := index; i < len(digits); i++ {
			digit := string(digits[index])
			letters := phoneMap[digit]
			for j := 0; j < len(letters); j++ {
				tmp += string(letters[j])
				backtrack(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}

	backtrack(0)
	return combinations
}
