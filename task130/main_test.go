package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSolve calls solve
func TestSolve(t *testing.T) {

	type solveTest struct {
		board  [][]byte
		result [][]byte
	}

	var solveTests = []solveTest{
		{
			board:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			result: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			board:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X'}},
			result: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'O', 'X'}},
		},
		{
			board:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'X', 'O', 'X'}},
			result: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'X', 'O', 'X'}},
		},
		{
			board:  [][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X', 'X'}, {'X', 'X', 'O', 'X', 'X'}, {'X', 'O', 'X', 'X', 'X'}},
			result: [][]byte{{'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X', 'X'}},
		},
		{
			board:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			result: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			board:  [][]byte{{'X'}},
			result: [][]byte{{'X'}},
		},
		{
			board:  [][]byte{{'X', 'O'}, {'X', 'O'}},
			result: [][]byte{{'X', 'O'}, {'X', 'O'}},
		},
		{
			board:  [][]byte{{'X', 'X'}, {'X', 'X'}},
			result: [][]byte{{'X', 'X'}, {'X', 'X'}},
		},
	}

	for _, test := range solveTests {
		solve(test.board)
		assert.Equal(t, test.result, test.board)
	}
}
