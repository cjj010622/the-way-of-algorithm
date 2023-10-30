package day53

import "sort"

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。

candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	var tmp []int

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			res = append(res, append([]int{}, tmp...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if candidates[i] > target-sum {
				break
			}

			tmp = append(tmp, candidates[i])
			backtrack(i, sum+candidates[i])
			tmp = tmp[:len(tmp)-1]
		}
	}

	backtrack(0, 0)
	return res
}
