package day52

import (
	"container/heap"
	"math"
	"sort"
)

/*
给你一个整数数组 gifts ，表示各堆礼物的数量。每一秒，你需要执行以下操作：

选择礼物数量最多的那一堆。
如果不止一堆都符合礼物数量最多，从中选择任一堆即可。
选中的那一堆留下平方根数量的礼物（向下取整），取走其他的礼物。
返回在 k 秒后剩下的礼物数量。
*/

func pickGifts(gifts []int, k int) int64 {
	h := &hp{gifts}
	heap.Init(h)
	for ; k > 0; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		heap.Fix(h, 0)
	}
	var ans int64
	for _, x := range gifts {
		ans += int64(x)
	}
	return ans
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h hp) Pop() (_ any) {
	return
}

func (h hp) Push(any) {
}
