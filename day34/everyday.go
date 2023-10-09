package day34

import (
	"sort"
	"strconv"
)

/*
给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：

num1 和 num2 直接连起来，得到 num 各数位的一个排列。
换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
num1 和 num2 可以包含前导 0 。
请你返回 num1 和 num2 可以得到的和的 最小 值。
*/

func splitNum(num int) int {
	numStr := []byte(strconv.Itoa(num))
	sort.Slice(numStr, func(i, j int) bool {
		return numStr[i] < numStr[j]
	})

	num1, num2 := 0, 0
	for i := 0; i < len(numStr); i++ {
		if i%2 == 0 {
			num1 = num1*10 + int(numStr[i]-'0')
		} else {
			num2 = num2*10 + int(numStr[i]-'0')
		}
	}

	return num1 + num2
}
