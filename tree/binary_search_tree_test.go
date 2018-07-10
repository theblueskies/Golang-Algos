package bst

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBSTInsert checks the Insert function
func TestBSTInsert(t *testing.T) {
	bst := BinarySearchTree{}
	v := []int{10, 7, 15, 5, 9, 11, 17}

	for _, num := range v {
		bst.Insert(num, strconv.Itoa(num))
	}

	assert.Equal(t, 10, bst.root.key)
	assert.Equal(t, "10", bst.root.value)
}
