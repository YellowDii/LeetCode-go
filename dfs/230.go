package dfs

// 230.二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历
	var res []int
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}
	inorder(root)
	return res[k-1]
}

// 迭代
func kthSmallest2(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

type MyBst struct {
	root    *TreeNode
	nodeNum map[*TreeNode]int // 统计以每个结点为根结点的子树的结点数，并存储在哈希表中
}

// 统计以 node 为根结点的子树的结点数
func (t *MyBst) countNodeNum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	t.nodeNum[node] = 1 + t.countNodeNum(node.Left) + t.countNodeNum(node.Right)
	return t.nodeNum[node]
}

// 返回二叉搜索树中第 k 小的元素
func (t *MyBst) kthSmallest(k int) int {
	node := t.root
	for {
		leftNodeNum := t.nodeNum[node.Left]
		if leftNodeNum < k-1 {
			node = node.Right
			k -= leftNodeNum + 1
		} else if leftNodeNum == k-1 {
			return node.Val
		} else {
			node = node.Left
		}
	}
}

// 记录以每个结点为根结点的子树的结点数
func kthSmallest3(root *TreeNode, k int) int {
	t := &MyBst{root, map[*TreeNode]int{}}
	t.countNodeNum(root)
	return t.kthSmallest(k)
}
