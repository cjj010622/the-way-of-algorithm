package day35

import "sort"

/*
有一些机器人分布在一条无限长的数轴上，他们初始坐标用一个下标从 0 开始的整数数组 nums 表示。当你给机器人下达命令时，它们以每秒钟一单位的速度开始移动。

给你一个字符串 s ，每个字符按顺序分别表示每个机器人移动的方向。'L' 表示机器人往左或者数轴的负方向移动，'R' 表示机器人往右或者数轴的正方向移动。

当两个机器人相撞时，它们开始沿着原本相反的方向移动。

请你返回指令重复执行 d 秒后，所有机器人之间两两距离之和。由于答案可能很大，请你将答案对 109 + 7 取余后返回。
*/

func sumDistance(nums []int, s string, d int) int {
	const mod int = 1e9 + 7
	for i, v := range s {
		if v == 'R' {
			nums[i] += d
		} else {
			nums[i] -= d
		}
	}

	sort.Ints(nums)
	var ans, tmp int
	for i, v := range nums {
		ans = (ans + i*v - tmp) % mod
		tmp += v
	}

	return ans
}
