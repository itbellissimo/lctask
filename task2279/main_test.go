package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMaximumBags calls maximumBags
func TestMaximumBags(t *testing.T) {

	type maximumBagsTest struct {
		capacity        []int
		rocks           []int
		additionalRocks int
		result          int
	}

	var maximumBagsTests = []maximumBagsTest{
		{
			capacity:        []int{2, 3, 4, 5},
			rocks:           []int{1, 2, 4, 4},
			additionalRocks: 2,
			result:          3,
		},
		{
			capacity:        []int{10, 2, 2},
			rocks:           []int{2, 2, 0},
			additionalRocks: 100,
			result:          3,
		},
	}

	for _, test := range maximumBagsTests {
		res := maximumBags(test.capacity, test.rocks, test.additionalRocks)
		assert.Equal(t, test.result, res)
	}
}
