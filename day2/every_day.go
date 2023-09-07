package day2

import "math"

// 修车的最少时间 二分查找最小时间

func repairCars(ranks []int, cars int) int64 {
	l, r := 1, ranks[0]*cars*cars
	check := func(m int) bool {
		cnt := 0
		for _, v := range ranks {
			cnt += int(math.Sqrt(float64(m / v)))
		}
		return cnt >= cars
	}

	for l < r {
		m := (l + r) >> 1
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return int64(l)
}
