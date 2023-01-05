package main

import (
	"github.com/itbellissimo/treenode"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCountSubIslands calls countSubIslands
func TestFindTilt(t *testing.T) {

	type findTiltTest struct {
		input  string
		result int
		err    error
	}

	var findTiltTests = []findTiltTest{
		{
			input:  "1,2,3",
			result: 1,
		},
		{
			input:  "4,2,9,3,5,null,7",
			result: 15,
		},
		{
			input:  "21,7,14,1,1,2,2,3,3",
			result: 9,
		},
	}

	for _, test := range findTiltTests {
		treeStr := treenode.NewTreeString(test.input)
		root, err := treeStr.GetTreeNode()
		if err != nil && test.err == nil {
			assert.Failf(t, "error", "error %v, %v", err, test.err)
			return
		}

		res := findTilt(root)
		assert.Equal(t, test.result, res)
	}
}
