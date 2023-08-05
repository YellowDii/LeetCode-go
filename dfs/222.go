package dfs

import "sort"

// 222.完全二叉树的节点个数
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}

	// 二分查找，范围[2^level, 2^(level+1)-1]
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			// 比min还小
			return false
		}
		bits := 1 << (level - 1) // 根为1 从根下第一层层开始
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				// K的高度对应二进制位为0，往左走
				node = node.Left
			} else {
				// K的当前高度对应二进制位为1，往右走
				node = node.Right
			}
			bits >>= 1 // 下一层
		}
		return node == nil
	}) - 1 // search返回的是实际个数+1的那个位置
}
