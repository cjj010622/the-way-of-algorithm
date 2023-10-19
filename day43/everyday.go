package day43

/*
给你一个由 不同 正整数组成的数组 nums ，请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d 。
*/

func tupleSameProduct(nums []int) int {
	cnt:=map[int]int{}
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			cnt[nums[i]*nums[j]]++
		}
	}
	ans:=0
	for _,v:=range cnt{
		ans+=v*(v-1)*4
	}
	return ans
}