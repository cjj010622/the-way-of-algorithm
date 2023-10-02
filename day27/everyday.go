package day27

/*
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。
*/

func maxProfit(prices []int) int {
    dp0,dp1:=0,-prices[0]
    for i:=1;i<len(prices);i++{
        dp0,dp1=max(dp0,dp1+prices[i]),max(dp1,dp0-prices[i])
    }

    return dp0
}

func max (a,b int) int{
    if a>b {
        return a
    }

    return b
}