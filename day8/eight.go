package day8

//买卖股票最佳时机2

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return prices[0]
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
