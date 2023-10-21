package day45

/*
给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。

请你返回 无法互相到达 的不同 点对数目 。
*/

func countPairs(n int, edges [][]int) int64 {
	g := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(i int) int {
		if vis[i] {
			return 0
		}
		vis[i] = true
		cnt := 1
		for _, j := range g[i] {
			cnt += dfs(j)
		}
		return cnt
	}
	var s, ans int64
	for i := 0; i < n; i++ {
		t := int64(dfs(i))
		ans += s * t
		s += t
	}
	return ans
}
