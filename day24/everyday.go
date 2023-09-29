package day24

/*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	m := len(flowerbed)
	for i := 0; i < m; i++ {
		if (i == 0 || flowerbed[i-1] == 0) && (flowerbed[i] == 0 && (i == m-1 || flowerbed[i+1] == 0)) {
			n--
			i++
		}
	}
	return n <= 0
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}