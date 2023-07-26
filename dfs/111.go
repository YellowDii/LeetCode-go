package dfs

// 111.二叉树的最小深度
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		if root.Left == nil && root.Right == nil {
			return 1
		} else if root.Left != nil && root.Right == nil {
			return 1 + minDepth(root.Left)
		} else if root.Left == nil && root.Right != nil {
			return 1 + minDepth(root.Right)
		} else {
			return 1 + min(minDepth(root.Left), minDepth(root.Right))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
