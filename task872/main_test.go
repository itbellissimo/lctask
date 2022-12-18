package main

import (
	"github.com/itbellissimo/treenode"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConstructRectangle calls updateMatrix
func TestUpdateMatrix(t *testing.T) {

	type leafSimilarTest struct {
		tree1  string
		tree2  string
		result bool
	}

	var leafSimilarTests = []leafSimilarTest{
		{
			tree1:  "3,5,1,6,2,9,8,null,null,7,4",
			tree2:  "3,5,1,6,7,4,2,null,null,null,null,null,null,9,8",
			result: true,
		},
		{
			tree1:  "1,2,3",
			tree2:  "1,3,2",
			result: false,
		},
	}

	for _, test := range leafSimilarTests {
		root1, err := treenode.NewTreeString(test.tree1).GetTreeNode()
		if err != nil {
			assert.Fail(t, err.Error())
		}
		root2, err := treenode.NewTreeString(test.tree2).GetTreeNode()
		if err != nil {
			assert.Fail(t, err.Error())
		}

		res := leafSimilar(root1, root2)
		assert.Equal(t, res, test.result)
	}
}
