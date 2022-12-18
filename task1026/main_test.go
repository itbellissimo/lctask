package main

import (
	"github.com/itbellissimo/treenode"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMaxAncestorDiff calls maxAncestorDiff
func TestMaxAncestorDiff(t *testing.T) {

	type maxAncestorDiffTest struct {
		tree   string
		result int
		err    error
	}

	var maxAncestorDiffTests = []maxAncestorDiffTest{
		{
			tree:   "8,3,10,1,6,null,14,null,null,4,7,13",
			result: 7,
		},
		{
			tree:   "1,null,2,null,0,3",
			result: 3,
		},
	}

	for _, test := range maxAncestorDiffTests {
		node, err := treenode.NewTreeString(test.tree).GetTreeNode()
		if err != nil && test.err == nil {
			assert.Failf(t, "error", "error %v, %v", err, test.err)
			return
		}
		res := maxAncestorDiff(node)
		assert.Equal(t, res, test.result)
	}
}
