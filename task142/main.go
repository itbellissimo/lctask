package main

import (
	"fmt"
)

func main() {
	//Input: head = [3,2,0,-4], pos = 1
	//Output: tail connects to node index 1
	//Explanation: There is a cycle in the linked list, where tail connects to the second node.
	var head *ListNode
	head = &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	head.Next.Next.Next.Next = head.Next

	fmt.Println(head)
	res := detectCycle(head)
	fmt.Println(res)
	i := detectCycleIndex(head)
	fmt.Println(i)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	i := 0
	index := map[*ListNode]int{head: i}

	for head.Next != nil {
		if _, ok := index[head.Next]; ok {
			return head.Next
		}

		i++
		index[head.Next] = i
	}

	return nil
}

// Maybe function in task is incorrect.
// According to task description we should return integer index position.

func detectCycleIndex(head *ListNode) int {
	count := 0
	index := map[*ListNode]int{head: count}

	for head.Next != nil {
		if i, ok := index[head.Next]; ok {
			return i
		}

		count++
		index[head.Next] = count
		head = head.Next
	}

	return -1
}
