package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMoveZeroes calls moveZeroes
func TestMoveZeroes(t *testing.T) {

	type moveZeroesTest struct {
		input  []int
		output []int
	}

	var moveZeroesTests = []moveZeroesTest{
		{
			input:  []int{0, 1, 0, 3, 12},
			output: []int{1, 3, 12, 0, 0},
		},
		{
			input:  []int{0},
			output: []int{0},
		},
		{
			input:  []int{0, 0, 0},
			output: []int{0, 0, 0},
		},
		{
			input:  []int{0, 1, 0, 3, 12, 0, 0, 0, 0},
			output: []int{1, 3, 12, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, test := range moveZeroesTests {
		moveZeroes(test.input)
		assert.Equal(t, test.output, test.input)
	}
}
