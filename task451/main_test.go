package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestFrequencySort calls frequencySort
func TestFrequencySort(t *testing.T) {

	type frequencySortTest struct {
		s      string
		result []string
	}

	var frequencySortTests = []frequencySortTest{
		{
			s:      "tree",
			result: []string{"eert", "eetr"},
		},
		{
			s:      "cccaaa",
			result: []string{"cccaaa", "aaaccc"},
		},
		{
			s:      "Aabb",
			result: []string{"bbAa"},
		},
	}

	for _, test := range frequencySortTests {
		res := frequencySort(test.s)
		assert.Contains(t, test.result, res)
	}
}
