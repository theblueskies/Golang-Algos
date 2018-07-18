package stringops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompress(t *testing.T) {
	c := Compress("AAABBBCDZZZZ")
	assert.Equal(t, "A2B3C1D1Z4", c)
}
