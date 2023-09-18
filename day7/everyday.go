package day7

/*
你总共需要上 numCourses 门课，课程编号依次为 0 到 numCourses-1 。你会得到一个数组 prerequisite ，其中 prerequisites[i] = [ai, bi] 表示如果你想选 bi 课程，你 必须 先选 ai 课程。

有的课会有直接的先修课程，比如如果想上课程 1 ，你必须先上课程 0 ，那么会以 [0,1] 数对的形式给出先修课程数对。
先决条件也可以是 间接 的。如果课程 a 是课程 b 的先决条件，课程 b 是课程 c 的先决条件，那么课程 a 就是课程 c 的先决条件。

你也得到一个数组 queries ，其中 queries[j] = [uj, vj]。对于第 j 个查询，您应该回答课程 uj 是否是课程 vj 的先决条件。

返回一个布尔数组 answer ，其中 answer[j] 是第 j 个查询的答案。
*/

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	var (
		edges = make([][]int, numCourses)
		isPre = make([][]bool, numCourses)
		indeg = make([]int, numCourses)
	)
	for i, _ := range isPre {
		isPre[i] = make([]bool, numCourses)
	}

	for _, info := range prerequisites {
		edges[info[0]] = append(edges[info[0]], info[1])
		indeg[info[1]]++
	}

	var q []int
	for k, v := range indeg {
		if v == 0 {
			q = append(q, k)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, v := range edges[cur] {
			isPre[cur][v] = true
			for i := 0; i < numCourses; i++ {
				isPre[i][v] = isPre[i][v] || isPre[i][cur]
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	var res []bool
	for _, query := range queries {
		res = append(res, isPre[query[0]][query[1]])
	}

	return res
}
