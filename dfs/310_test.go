package dfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	edges := [][]int{{1, 0}, {1, 2}, {1, 3}}
	assert.Equal(t, []int{1}, findMinHeightTrees(4, edges))
}
