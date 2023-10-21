package day45

/*
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
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

	return len(res) == numCourses
}