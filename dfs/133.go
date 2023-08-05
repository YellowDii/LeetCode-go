package dfs

// 133.克隆图

// Node133 Definition .
type Node133 struct {
	Val       int
	Neighbors []*Node133
}

func cloneGraph(node *Node133) *Node133 {
	if node == nil {
		return node
	}
	visited := make(map[*Node133]*Node133)
	var dfs func(*Node133) *Node133
	dfs = func(node *Node133) *Node133 {
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		cloneNode := &Node133{Val: node.Val}
		visited[node] = cloneNode
		for _, neighbor := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(neighbor))
		}
		return cloneNode
	}
	return dfs(node)
}
