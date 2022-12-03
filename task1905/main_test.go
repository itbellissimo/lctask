package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCountSubIslands calls countSubIslands
func TestCountSubIslands(t *testing.T) {

	type countSubIslandsTest struct {
		grid1  [][]int
		grid2  [][]int
		result int
	}

	var countSubIslandsTests = []countSubIslandsTest{
		{
			grid1:  [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			grid2:  [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}},
			result: 3,
		},
		{
			grid1:  [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
			grid2:  [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}},
			result: 2,
		},
	}

	for _, test := range countSubIslandsTests {
		res := countSubIslands(test.grid1, test.grid2)
		assert.Equal(t, test.result, res)
	}
}
