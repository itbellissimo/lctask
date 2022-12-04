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
	floydAlgoRes := detectCycleFloyds(head)
	fmt.Println(floydAlgoRes)
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

// https://leetcode.com/problems/linked-list-cycle-ii/discuss/726011/A-mechanics-based-visualization-(%2B-Floyds-algo-explained)
func detectCycleFloyds(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast != slow || head == nil || head.Next == nil {
		return nil
	}
	count := 1
	curr := fast.Next
	for curr != fast {
		curr = curr.Next
		count++
	}
	//count === len of the loop
	front, rear := head, head
	for i := 0; i < count; i++ {
		front = front.Next //get a distance of L between them
	}
	for front != rear { //keep moving both at same rate until they meet
		front = front.Next
		rear = rear.Next
	}
	return rear //or front

}
