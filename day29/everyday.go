package day29

/*
给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit(k int, prices []int) int {
    f:=make([][2]int,k+1)
    for j:=1;j<=k;j++{
        f[j][1]=-prices[0]
    }

    for _,x:=range prices[1:]{
        for j:=k;j>0;j--{
            f[j][0]=max(f[j][0],f[j][1]+x)
            f[j][1]=max(f[j][1],f[j-1][0]-x)
        }
    }
    return f[k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}