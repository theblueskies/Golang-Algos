package stringops

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalancedBrackets(t *testing.T) {
	var flagtests = []struct {
		input  string
		output string
	}{
		{"{}", "Valid"},
		{"{}{}", "Valid"},
		{"{}}", "Invalid"},
		{"{", "Invalid"},
		{"{()", "Invalid"},
		{"{(})", "Invalid"},
		{"(}{})", "Invalid"},
		{"{(({}))}", "Valid"},
	}
	for _, tt := range flagtests {
		r := CheckValidBalance(tt.input)
		assert.Equal(t, tt.output, r)
	}
}
