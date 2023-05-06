package day1

import "math"

// 凑零钱，暴力即遍历n叉树
func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	res := math.MinInt32

	//暴力遍历每个硬币
	for _, coin := range coins {
		subProblem := dp(coins, amount-coin)

		//如果当前硬币无解就跳过
		if subProblem == -1 {
			continue
		}

		//在子问题中选择最优解然后加1
		res = int(math.Min(float64(res), float64(subProblem+1)))
	}

	//如果未发送变化说明无解
	if res == math.MinInt32 {
		return -1
	}

	return res
}
