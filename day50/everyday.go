package day50

/*
给你一个整数 num ，返回 num 中能整除 num 的数位的数目。

如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums 。
*/

func countDigits(num int) int {
	tmp := num
	var ans, t int
	for tmp > 0 {
		t = tmp % 10
		tmp = tmp / 10
		if num%t == 0 {
			ans++
		}
	}
	return ans
}
