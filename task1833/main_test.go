package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestMaxIceCream calls maxIceCream
func TestMaxIceCream(t *testing.T) {

	type maxIceCreamTest struct {
		cost   []int
		coins  int
		result int
	}

	var maxIceCreamTests = []maxIceCreamTest{
		{
			cost:   []int{1, 3, 2, 4, 1},
			coins:  7,
			result: 4,
		},
		{
			cost:   []int{10, 6, 8, 7, 7, 8},
			coins:  5,
			result: 0,
		},
		{
			cost:   []int{1, 6, 3, 1, 2, 5},
			coins:  20,
			result: 6,
		},
	}

	for _, test := range maxIceCreamTests {
		res := maxIceCream(test.cost, test.coins)
		assert.Equal(t, test.result, res)
	}
}
