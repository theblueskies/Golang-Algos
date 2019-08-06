package diagonalsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalSum(t *testing.T) {
	arr := [][]int32{
		[]int32{3},
		[]int32{11, 2, 4},
		[]int32{4, 5, 6},
		[]int32{10, 8, -12},
	}

	result := diagonalDifference(arr)
	assert.Equal(t, int32(15), result)
}
