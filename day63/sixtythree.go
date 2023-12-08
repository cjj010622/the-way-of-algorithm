package day63

import "fmt"

/*
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
*/

func addBinary(a string, b string) string {
	var result string
	carry := 0

	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		result = fmt.Sprintf("%d%s", sum%2, result)
		carry = sum / 2
	}

	return result
}
