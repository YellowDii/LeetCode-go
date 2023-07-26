package dfs

var localHigh = make(map[*TreeNode]int)

// 110.平衡二叉树
func isBalanced(root *TreeNode) bool {
	// 方法1
	if root == nil {
		return true
	}

	if isBalanced(root.Left) && isBalanced(root.Right) && abs(high(root.Left)-high(root.Right)) <= 1 {
		return true
	} else {
		return false
	}
	// 方法2
	// _, ok := isBalancedHelper(root)
	// return ok
}

func high(root *TreeNode) int {
	if root == nil {
		localHigh[root] = 0
		return 0
	} else {
		if val, ok := localHigh[root]; ok {
			return val
		} else {
			left := high(root.Left)
			right := high(root.Right)
			localHigh[root] = 1 + max(left, right)
			return localHigh[root]
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalancedHelper(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, ok := isBalancedHelper(root.Left)
	if !ok {
		return 0, false
	}
	rightDepth, ok := isBalancedHelper(root.Right)
	if !ok {
		return 0, false
	}

	if leftDepth >= rightDepth {
		return leftDepth + 1, leftDepth-rightDepth <= 1
	} else {
		return rightDepth + 1, rightDepth-leftDepth <= 1
	}

}
