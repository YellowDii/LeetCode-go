package dfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116.填充每个节点的下一个右侧节点指针
func connect116(root *Node) *Node {
	if root == nil {
		return root
	}

	root.Left = connect116(root.Left)
	root.Right = connect116(root.Right)
	left := root.Left
	right := root.Right
	for left != nil && right != nil {
		left.Next = right
		left = left.Right
		right = right.Left
	}

	return root
}
