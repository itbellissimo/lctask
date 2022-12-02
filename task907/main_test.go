package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSumSubarrayMins calls sumSubarrayMins
func TestSumSubarrayMins(t *testing.T) {

	type sumSubarrayMinsTest struct {
		arr    []int
		result int
	}

	var sumSubarrayMinsTests = []sumSubarrayMinsTest{
		{
			arr:    []int{0, 0, 0, 0, 0, 0, 0},
			result: 0,
		},
		{
			arr:    []int{3, 1, 2, 4},
			result: 17,
		},
		{
			arr:    []int{11, 81, 94, 43, 3},
			result: 444,
		},
		{
			arr:    []int{0, 0, 1, -7, 77, 33, 11, 0, 1, 100},
			result: 83,
		},
	}

	for _, test := range sumSubarrayMinsTests {
		res := sumSubarrayMins(test.arr)
		assert.Equal(t, test.result, res)
	}
}
