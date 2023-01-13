package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLongestPath calls longestPath
func TestLongestPath(t *testing.T) {

	type longestPathTest struct {
		parent []int
		s      string
		result int
	}

	var longestPathTests = []longestPathTest{
		{
			parent: []int{-1, 0, 0, 1, 1, 2},
			s:      "abacbe",
			result: 3,
		},
		{
			parent: []int{-1, 0, 0, 0},
			s:      "aabc",
			result: 3,
		},
	}

	for _, test := range longestPathTests {
		res := longestPath(test.parent, test.s)
		assert.Equal(t, test.result, res)
	}
}
