package day49

import "strconv"

/*
给你一个正整数 n ，请你返回 n 的 惩罚数 。

n 的 惩罚数 定义为所有满足以下条件 i 的数的平方和：

1 <= i <= n
i * i 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 i 。
*/

func punishmentNumber(n int) int {
	var check func(string,int,int) bool
	check=func(s string,i,x int) bool{
		m:=len(s)
		if i>=m{
			return x==0
		}
		y:=0
		for j:=i;j<m;j++{
			y=y*10+int(s[j]-'0')
			if y>x{
				break
			}
			if check(s,j+1,x-y){
				return true
			}
		}
		return false
	}
	var ans int
	for i:=1;i<=n;i++{
		x:=i*i
		s:=strconv.Itoa(x)
		if check(s,0,i){
			ans += x
		}

	}
	return ans
}