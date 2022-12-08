package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// TestConstructRectangle calls updateMatrix
func TestUpdateMatrix(t *testing.T) {

	type updateMatrixTest struct {
		mat    [][]int
		result [][]int
	}

	var updateMatrixTests = []updateMatrixTest{
		{
			mat:    [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			result: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{1, 2, 1, 1, 1},
				{0, 1, 2, 2, 2},
				{1, 2, 3, 3, 3},
				{2, 3, 4, 4, 4},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 1, 2, 3, 4},
				{1, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6, 7},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 1, 2, 1, 0},
				{1, 2, 3, 2, 1},
				{2, 3, 4, 3, 2},
				{3, 4, 5, 4, 3},
			},
		},
		{
			mat: [][]int{
				{0},
				{1},
				{1},
				{1},
			},
			result: [][]int{
				{0},
				{1},
				{2},
				{3},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 0},
			},
			result: [][]int{
				{0, 1, 2, 1, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{1, 1, 1, 1, 1},
			},
		},
		{
			mat: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
		},
		{
			mat: [][]int{
				{},
			},
			result: [][]int{
				{},
			},
		},
		{
			mat: [][]int{
				{},
				{},
			},
			result: [][]int{
				{},
				{},
			},
		},
	}

	for _, test := range updateMatrixTests {
		res := updateMatrix(test.mat)
		assert.Equal(t, len(test.mat[0]), len(test.result[0]))
		assert.Equal(t, len(test.mat), len(test.result))
		for i := 0; i < len(test.result); i++ {
			for j := 0; j < len(test.result[0]); j++ {
				assert.Equal(t, test.result[i][j], res[i][j])
			}
		}
	}
}

// TestConstructRectangle calls updateMatrixFromSite
func TestUpdateMatrixFromSite(t *testing.T) {

	type updateMatrixFromSiteTest struct {
		mat    [][]int
		result [][]int
	}

	var updateMatrixFromSiteTests = []updateMatrixFromSiteTest{
		{
			mat:    [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			result: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0},
				{1, 2, 1, 1, 1},
				{0, 1, 2, 2, 2},
				{1, 2, 3, 3, 3},
				{2, 3, 4, 4, 4},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 1, 2, 3, 4},
				{1, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6, 7},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 0},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{0, 1, 2, 1, 0},
				{1, 2, 3, 2, 1},
				{2, 3, 4, 3, 2},
				{3, 4, 5, 4, 3},
			},
		},
		{
			mat: [][]int{
				{0},
				{1},
				{1},
				{1},
			},
			result: [][]int{
				{0},
				{1},
				{2},
				{3},
			},
		},
		{
			mat: [][]int{
				{0, 1, 1, 1, 0},
			},
			result: [][]int{
				{0, 1, 2, 1, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			result: [][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
		},
		{
			mat: [][]int{
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt},
			},
		},
		{
			mat: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			result: [][]int{
				{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt},
				{math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt},
			},
		},
		{
			mat: [][]int{
				{},
			},
			result: [][]int{
				{},
			},
		},
		{
			mat: [][]int{
				{},
				{},
			},
			result: [][]int{
				{},
				{},
			},
		},
	}

	for _, test := range updateMatrixFromSiteTests {
		res := updateMatrixFromSite(test.mat)
		assert.Equal(t, len(test.mat[0]), len(test.result[0]))
		assert.Equal(t, len(test.mat), len(test.result))
		for i := 0; i < len(test.result); i++ {
			for j := 0; j < len(test.result[0]); j++ {
				assert.Equal(t, test.result[i][j], res[i][j])
			}
		}
	}
}
