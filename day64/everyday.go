package day64

import "sort"

/*
你驾驶出租车行驶在一条有 n 个地点的路上。这 n 个地点从近到远编号为 1 到 n ，你想要从 1 开到 n ，通过接乘客订单盈利。你只能沿着编号递增的方向前进，不能改变方向。

乘客信息用一个下标从 0 开始的二维数组 rides 表示，其中 rides[i] = [starti, endi, tipi] 表示第 i 位乘客需要从地点 starti 前往 endi ，愿意支付 tipi 元的小费。

每一位 你选择接单的乘客 i ，你可以 盈利 endi - starti + tipi 元。你同时 最多 只能接一个订单。

给你 n 和 rides ，请你返回在最优接单方案下，你能盈利 最多 多少元。

注意：你可以在一个地点放下一位乘客，并在同一个地点接上另一位乘客。
*/

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool { return rides[i][0] < rides[j][0] })
	m := len(rides)
	f := make([]int64, m)
	var dfs func(int) int64
	dfs = func(i int) int64 {
		if i >= m {
			return 0
		}

		if f[i] == 0 {
			st, ed, tip := rides[i][0], rides[i][1], rides[i][2]
			j := sort.Search(m, func(j int) bool { return rides[j][0] >= ed })
			f[i] = max(dfs(i+1), dfs(j)+int64(ed-st+tip))
		}
		return f[i]
	}
	return dfs(0)
}
