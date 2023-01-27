package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMinDeletionSize calls minDeletionSize
func TestMinDeletionSize(t *testing.T) {

	type minDeletionSizeTest struct {
		strs   []string
		result int
	}

	var minDeletionSizeTests = []minDeletionSizeTest{
		{
			strs:   []string{"abc", "bce", "cae"},
			result: 1,
		},
		{
			strs:   []string{"cba", "daf", "ghi"},
			result: 1,
		},
		{
			strs:   []string{"a", "b"},
			result: 0,
		},
		{
			strs:   []string{"zyx", "wvu", "tsr"},
			result: 3,
		},
	}

	for _, test := range minDeletionSizeTests {
		res := minDeletionSize(test.strs)
		assert.Equal(t, test.result, res)
	}
}
