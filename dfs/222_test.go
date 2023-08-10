package dfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countNodes2(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{
		Val: 2,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 5,
	}
	root.Right.Left = &TreeNode{Val: 6}
	assert.Equal(t, 6, countNodes2(root))
}
