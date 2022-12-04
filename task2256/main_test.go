package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSolve calls solve
func TestMinimumAverageDifference(t *testing.T) {

	type minimumAverageDifferenceTest struct {
		nums   []int
		result int
	}

	var minimumAverageDifferenceTests = []minimumAverageDifferenceTest{
		{
			nums:   []int{2, 5, 3, 9, 5, 3},
			result: 3,
		},
		{
			nums:   []int{0},
			result: 0,
		},
	}

	for _, test := range minimumAverageDifferenceTests {
		res := minimumAverageDifference(test.nums)
		assert.Equal(t, test.result, res)
	}
}
