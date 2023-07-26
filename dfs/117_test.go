package dfs

import "testing"

func Test_connect117(t *testing.T) {
	root := &Node{
		Val: 1,
	}
	root.Left = &Node{
		Val:   2,
		Left:  &Node{4, nil, nil, nil},
		Right: &Node{5, nil, nil, nil},
	}
	root.Right = &Node{
		Val:   3,
		Right: &Node{7, nil, nil, nil},
	}
	connect117(root)
}
