package day38

import "sort"

/*
你的国家有无数个湖泊，所有湖泊一开始都是空的。当第 n 个湖泊下雨前是空的，那么它就会装满水。如果第 n 个湖泊下雨前是 满的 ，这个湖泊会发生 洪水 。你的目标是避免任意一个湖泊发生洪水。

给你一个整数数组 rains ，其中：

rains[i] > 0 表示第 i 天时，第 rains[i] 个湖泊会下雨。
rains[i] == 0 表示第 i 天没有湖泊会下雨，你可以选择 一个 湖泊并 抽干 这个湖泊的水。
*/

func avoidFlood(rains []int) []int {
	n := len(rains)
	rainy := make(map[int]int)
	var sunny []int
	ans := make([]int, len(rains))
	for i := 0; i < n; i++ {
		ans[i] = -1
	}

	for i, v := range rains {
		if v > 0 {
			if j, ok := rainy[v]; ok {
				idx := sort.Search(len(sunny), func(i int) bool {
					return sunny[i] > j
				})
				if idx == len(sunny) {
					return []int{}
				}
				ans[sunny[idx]] = v
				sunny = append(sunny[:idx], sunny[idx+1:]...)
			}
			rainy[v] = i
		} else {
			sunny = append(sunny, i)
			ans[i] = 1
		}
	}
	return ans
}
