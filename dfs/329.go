package dfs

//329. 矩阵中的最长递增路径
var (
	dirs          = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	raws, columns int
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return 0
	}
	raws, columns = len(matrix), len(matrix[0])
	memo := make([][]int, raws)
	for i, _ := range memo {
		memo[i] = make([]int, columns)
	}
	ans := 0
	for i := 0; i < raws; i++ {
		for j := 0; j < columns; j++ {
			ans = max(ans, dfs(matrix, i, j, memo))
		}
	}
	return ans
}

func dfs(matrix [][]int, x int, y int, memo [][]int) int {
	if memo[x][y] != 0 {
		return memo[x][y]
	}
	memo[x][y]++
	for _, dir := range dirs {
		newRaw, newCol := x+dir[0], y+dir[1]
		if newRaw >= 0 && newRaw < raws && newCol >= 0 && newCol < columns && matrix[newRaw][newCol] > matrix[x][y] {
			memo[x][y] = max(memo[x][y], dfs(matrix, newRaw, newCol, memo)+1)
		}
	}
	return memo[x][y]
}
