package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestDetectCycleIndex calls detectCycleIndex
func TestDetectCycleIndex(t *testing.T) {

	type detectCycleIndexTest struct {
		input  *ListNode
		result int
	}

	var head *ListNode
	head = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: head,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next

	var head2 *ListNode
	head2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: head2,
		},
	}
	head2.Next.Next = head2

	var head3 *ListNode
	head3 = &ListNode{
		Val:  1,
		Next: nil,
	}

	var detectCycleIndexTests = []detectCycleIndexTest{
		{
			input:  head,
			result: 1,
		},
		{
			input:  head2,
			result: 0,
		},
		{
			input:  head3,
			result: -1,
		},
	}

	for _, test := range detectCycleIndexTests {
		res := detectCycleIndex(test.input)
		assert.Equal(t, test.result, res)
	}
}
