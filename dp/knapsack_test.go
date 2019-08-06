package dp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnapSackItem(t *testing.T) {
	i := []KnapSackItem{
		KnapSackItem{Weight: 10, value: 5},
		KnapSackItem{Weight: 11, value: 7},
		KnapSackItem{},
	}

	for _, v := range i {
		fmt.Println(v)
	}

	assert.NotEmpty(t, i)

}

func TestGetValue(t *testing.T) {
	i := []KnapSackItem{
		KnapSackItem{Weight: 10, value: 5},
		KnapSackItem{Weight: 11, value: 7},
		KnapSackItem{},
	}

	v := GetValue(i)
	assert.Equal(t, 12, v)
}

func TestGetMostValuableItems(t *testing.T) {
	arr := []KnapSackItem{
		KnapSackItem{Weight: 10, value: 5},
		KnapSackItem{Weight: 11, value: 7},
		KnapSackItem{Weight: 2, value: 4},
		KnapSackItem{Weight: 5, value: 1},
		KnapSackItem{Weight: 7, value: 6},
		KnapSackItem{Weight: 1, value: 4},
	}

	e := []KnapSackItem{
		KnapSackItem{Weight: 2, value: 4},
		KnapSackItem{Weight: 7, value: 6},
		KnapSackItem{Weight: 1, value: 4},
	}

	r := GetMostValuableItems(arr, []KnapSackItem{}, 12, 12)
	assert.Equal(t, e, r)
}
