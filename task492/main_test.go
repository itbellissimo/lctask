package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConstructRectangle calls constructRectangle
func TestConstructRectangle(t *testing.T) {

	type constructRectangleTest struct {
		area   int
		result []int
	}

	var constructRectangleTests = []constructRectangleTest{
		{
			area:   4,
			result: []int{2, 2},
		},
		{
			area:   37,
			result: []int{37, 1},
		},
		{
			area:   122122,
			result: []int{427, 286},
		},
	}

	for _, test := range constructRectangleTests {
		res := constructRectangle(test.area)
		assert.Equal(t, len(test.result), 2)
		assert.Equal(t, test.result[0], res[0])
		assert.Equal(t, test.result[1], res[1])
	}
}
