package day46

/*
现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。

例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		graph = make([][]int, numCourses)
		indeg = make([]int, numCourses)
		queue []int
		res   []int
	)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		indeg[prerequisite[0]]++
	}

	for k, v := range indeg {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		for _, info := range graph[cur] {
			indeg[info]--
			if indeg[info] == 0 {
				queue = append(queue, info)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}

	return res
}
