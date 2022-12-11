package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConstructRectangle calls updateMatrix
func TestUpdateMatrix(t *testing.T) {

	type leafSimilarTest struct {
		tree1  []NilInt
		tree2  []NilInt
		result bool
	}

	var leafSimilarTests = []leafSimilarTest{
		{
			tree1: []NilInt{
				NewInt(3),
				NewInt(5),
				NewInt(1),
				NewInt(6),
				NewInt(2),
				NewInt(9),
				NewInt(8),
				NewNil(),
				NewNil(),
				NewInt(7),
				NewInt(4),
			},
			tree2: []NilInt{
				NewInt(3),
				NewInt(5),
				NewInt(1),
				NewInt(6),
				NewInt(7),
				NewInt(4),
				NewInt(2),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewNil(),
				NewInt(9),
				NewInt(8),
			},
			result: true,
		},
		{
			tree1: []NilInt{
				NewInt(1),
				NewInt(2),
				NewInt(3),
			},
			tree2: []NilInt{
				NewInt(1),
				NewInt(3),
				NewInt(2),
			},
			result: false,
		},
	}

	for _, test := range leafSimilarTests {
		line := treeLine(test.tree1)
		line2 := treeLine(test.tree2)
		res := leafSimilar(line.treeNode(), line2.treeNode())
		assert.Equal(t, res, test.result)
	}
}
