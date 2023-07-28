package dfs

var m, n int

// 130.被围绕的区域
func solve130(board [][]byte) {
	m = len(board)
	n = len(board[0])
	// 从边界开始搜索
	for i := 0; i < m; i++ {
		dfs130(board, i, 0)
		dfs130(board, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs130(board, 0, j)
		dfs130(board, m-1, j)
	}
	// 遍历board，将O变成X，将V变成O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'V' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	return
}

func dfs130(board [][]byte, i, j int) {
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'V' // 标记visited
	dfs130(board, i+1, j)
	dfs130(board, i-1, j)
	dfs130(board, i, j+1)
	dfs130(board, i, j-1)
}
