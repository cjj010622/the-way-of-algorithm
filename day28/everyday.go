package day28

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(prices []int) int {
    f1,f2,f3,f4:=-prices[0],0,-prices[0],0
    for i:=1;i<len(prices);i++{
        f1=max(f1,-prices[i])
        f2=max(f2,f1+prices[i])
        f3=max(f3,f2-prices[i])
        f4=max(f4,f3+prices[i])
    }
    return f4
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}