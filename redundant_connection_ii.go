package gopractice

//https://leetcode.com/problems/redundant-connection-ii

func findRedundantDirectedConnection(edges [][]int) []int {
	n := len(edges)
	tag := []int{}
	visited := []bool{}
	children := [][]int{}
	parents := []int{}
	for i := 0; i <= n; i++ {
		tag = append(tag, i)
		visited = append(visited, false)
		children = append(children, []int{})
		parents = append(parents, i)
	}
	find := func(v int) int {
		for tag[v] != v {
			v = tag[v]
		}
		return v
	}
	union := func(v1, v2 int) {
		tag[find(v2)] = find(v1)
	}

	candidateMoreParents := 0
	candidateCycle := 0
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		children[from] = append(children[from], to)
		if parents[to] != to {
			candidateMoreParents = to
		}
		parents[to] = from
		if find(from) == find(to) {
			candidateCycle = to
		}
		union(from, to)
	}

	if candidateMoreParents == 0 {
		return []int{parents[candidateCycle], candidateCycle}
	}

	var dfs func(v int) (int, int)
	dfs = func(v int) (int, int) {
		visited[v] = true
		for _, child := range children[v] {
			if visited[child] {
				return v, child
			}
			from, to := dfs(child)
			if from != 0 {
				return from, to
			}
		}
		return 0, 0
	}
	from, to := dfs(candidateMoreParents)
	if from == 0 {
		return []int{parents[candidateMoreParents], candidateMoreParents}
	} else {
		return []int{from, to}
	}
}