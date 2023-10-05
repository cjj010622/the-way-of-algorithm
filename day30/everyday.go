package day30

/*
给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    n := len(prices)
    f0, f1, f2 := -prices[0], 0, 0
    for i := 1; i < n; i++ {
        newf0 := max(f0, f2 - prices[i])
        newf1 := f0 + prices[i]
        newf2 := max(f1, f2)
        f0, f1, f2 = newf0, newf1, newf2
    }
    return max(f1, f2)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
