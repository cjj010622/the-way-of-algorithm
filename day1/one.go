package day1

//合并两个有序数组

func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			res = append(res, nums2[p2:]...)
			break
		}
		if p2 == n {
			res = append(res, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			res = append(res, nums1[p1])
			p1++
		} else {
			res = append(res, nums2[p2])
			p2++
		}
	}
	copy(nums1, res)
}
