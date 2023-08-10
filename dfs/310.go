package dfs

// 310.最小高度树
// bfs
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	parents := make([]int, n)
	bfs := func(start int) (node int) {
		queue := []int{start}
		visited := make([]bool, n)
		visited[start] = true
		for len(queue) > 0 {
			node, queue = queue[0], queue[1:] // 不要用:= 会变为for循环内的局部变量
			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					parents[neighbor] = node
					queue = append(queue, neighbor)
				}
			}
		}
		// 队列全部出栈完 node就是最后一个出栈的元素 即距离start最远的节点
		return
	}

	x := bfs(0) // 找到离0最远的节点x
	y := bfs(x) // 找到离x最远的节点y

	// 从y开始回溯到x
	path := []int{}
	parents[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parents[y]
	}
	m := len(path)
	if m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	}
	return []int{path[m/2]}

}

// dfs
func findMinHeightTrees2(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	parents := make([]int, n)
	maxDepth, node := 0, -1
	var dfs func(int, int, int)
	dfs = func(start, parent, depth int) {
		if depth > maxDepth {
			maxDepth, node = depth, start
		}
		parents[start] = parent
		for _, neighbor := range graph[start] {
			if neighbor != parent {
				dfs(neighbor, start, depth+1)
			}
		}
	}

	dfs(0, -1, 1) // 找到离0最远的节点x
	maxDepth = 0
	dfs(node, -1, 1) // 找到离x最远的节点y

	path := []int{}
	for node != -1 {
		path = append(path, node)
		node = parents[node]
	}

	if m := len(path); m%2 == 0 {
		return []int{path[m/2-1], path[m/2]}
	} else {
		return []int{path[m/2]}
	}
}
