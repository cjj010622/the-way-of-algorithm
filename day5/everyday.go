package day5

/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges  = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
		q      []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	for k, v := range indeg {
		if v == 0 {
			q = append(q, k)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, info := range edges[u] {
			indeg[info]--
			if indeg[info] == 0 {
				q = append(q, info)
			}
		}
	}

	if len(result) != numCourses {
		return []int{}
	}

	return result
}
