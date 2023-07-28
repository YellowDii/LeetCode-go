package dfs

// 129.求根到叶子节点数字之和
func sumNumbers(root *TreeNode) int {
	return dfs129(root, 0)
}

func dfs129(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return dfs129(root.Left, sum) + dfs129(root.Right, sum)
}
