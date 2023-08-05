package dfs

// 210.课程表II

func findOrder210(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(int)
	)
	dfs = func(u int) {
		if visited[u] == 0 {
			visited[u] = 1 // 标记为搜索中
			for _, v := range edges[u] {
				dfs(v)
				if !valid {
					// 如果发现有环，就退出
					return
				}
			}
			visited[u] = 2             // 标记为已完成
			result = append(result, u) // 将节点入栈\
		} else if visited[u] == 1 {
			valid = false
			return
		}
	}
	for _, edge := range prerequisites {
		edges[edge[1]] = append(edges[edge[1]], edge[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return nil
	} else {
		// 出栈
		for i := 0; i < len(result)/2; i++ {
			result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
		}
		return result
	}
}
