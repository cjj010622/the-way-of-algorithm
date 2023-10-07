package day28

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/

func isValid(s string) bool {
	var stack []rune
	for _, l := range s {
		if len(stack) == 0 {
			stack = append(stack, l)
		} else {
			t := stack[len(stack)-1]
			if (l == ')' && t == '(') ||
				(l == ']' && t == '[') ||
				(l == '}' && t == '{') {
				stack = append(stack[:len(stack)-1])
			} else {
				stack = append(stack, l)
			}
		}
	}
	return len(stack) == 0
}
