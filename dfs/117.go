package dfs

func connect117(root *Node) *Node {
	cur := root
	//O1的话，保存上一层的链表
	for cur != nil {
		dummy := &Node{}
		pre := dummy
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}
