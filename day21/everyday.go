package day21

/*
n 个人站成一排，按从 1 到 n 编号。

最初，排在队首的第一个人拿着一个枕头。每秒钟，拿着枕头的人会将枕头传递给队伍中的下一个人。一旦枕头到达队首或队尾，传递方向就会改变，队伍会继续沿相反方向传递枕头。

例如，当枕头到达第 n 个人时，TA 会将枕头传递给第 n - 1 个人，然后传递给第 n - 2 个人，依此类推。
给你两个正整数 n 和 time ，返回 time 秒后拿着枕头的人的编号。

*/

func passThePillow(n int, time int) int {
	d := 1
	h := 1
	for i := 1; i <= time; i++ {
		if h == 1 {
			d = 1
		} else if h == n {
			d = -1
		}
		h += d
	}
	return h
}
