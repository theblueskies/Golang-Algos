package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBSTInsert checks the Insert function
func TestBSTPathSearch(t *testing.T) {
	v := []int{10, 7, 15, 5, 9, 11, 17}
	ep := [][]int{
		[]int{10, 7},
		[]int{17},
	}

	paths := SearchPaths(v, 17)
	assert.Equal(t, ep, paths)
}

func TestEmptyBSTSearch(t *testing.T) {
	v := []int{}
	paths := SearchPaths(v, 17)
	assert.Equal(t, 0, len(paths))
}
