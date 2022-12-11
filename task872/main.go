package main

import (
	"fmt"
	"math"
)

type treeLine []NilInt

// [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8] [0:1] 0
// [5,1,6,7,4,2,null,null,null,null,null,null,9,8] [0:2] 1
// [6,7,4,2,null,null,null,null,null,null,9,8] [0:4] 2
// [null,null,null,null,null,null,9,8] [0:8] 3
//
//			          3
//		  	    5           1
//		  6         7    4        2
//	                            9    8

func (tl *treeLine) treeNode() *TreeNode {
	if (*tl)[0].null {
		fmt.Println("wrong root value", (*tl)[0].Value())
		return nil
	}

	lvl := 1
	tmp := (*tl)[1:]

	byLvl := make(map[int][]NilInt)

	byLvl[0] = make([]NilInt, 1, 1)
	byLvl[0][0] = (*tl)[0]

	for len(tmp) != 0 {
		end := int(math.Pow(2.0, float64(lvl)))
		if end >= len(tmp) {
			end = len(tmp)
		}
		values := tmp[:end]
		tmp = tmp[end:]

		byLvl[lvl] = make([]NilInt, 0, len(values))
		for i := 0; i < len(values); i = i + 2 {
			byLvl[lvl] = append(byLvl[lvl], values[i], values[i+1])
		}

		lvl++
	}

	lvl = len(byLvl) - 1
	nodeByLvl := make(map[int][]*TreeNode)

	for n := lvl; n >= 0; n-- {
		nodeByLvl[n] = make([]*TreeNode, len(byLvl[n]), len(byLvl[n]))
		for index, v := range byLvl[n] {
			var (
				l *TreeNode
				r *TreeNode
			)

			if n < lvl {
				if len(nodeByLvl[n+1])-1 >= index*2 && nodeByLvl[n+1][index*2] != nil {
					l = nodeByLvl[n+1][index*2]
				}
				if len(nodeByLvl[n+1])-1 >= index*2+1 && nodeByLvl[n+1][index*2+1] != nil {
					r = nodeByLvl[n+1][index*2+1]
				}
			}

			if v.null {
				nodeByLvl[n][index] = nil
			} else {
				nodeByLvl[n][index] = &TreeNode{
					Val:   v.value,
					Left:  l,
					Right: r,
				}
			}
		}
	}

	return nodeByLvl[0][0]
}

func main() {
	var (
		//root1 = []int{3, 5, 1, 6, 2, 9, 8, 7, 4}
		//root2 = []int{3, 5, 1, 6, 7, 4, 2, 9, 8}
		r1 = []NilInt{
			NewInt(3),
			NewInt(5),
			NewInt(1),
			NewInt(6),
			NewInt(2),
			NewInt(9),
			NewInt(8),
			NewNil(),
			NewNil(),
			NewInt(7),
			NewInt(4),
		}

		r2 = []NilInt{
			NewInt(3),
			NewInt(5),
			NewInt(1),
			NewInt(6),
			NewInt(7),
			NewInt(4),
			NewInt(2),
			NewNil(),
			NewNil(),
			NewNil(),
			NewNil(),
			NewNil(),
			NewNil(),
			NewInt(9),
			NewInt(8),
		}
	)

	fmt.Println(r1)
	fmt.Println(r2)

	line := treeLine(r1)
	fmt.Println(line.treeNode())

	line2 := treeLine(r2)
	fmt.Println(line2.treeNode())

	mat2 := leafSimilar(line.treeNode(), line2.treeNode())
	fmt.Println(mat2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func walk(node *TreeNode, ch chan int) {
	if node == nil {
		return
	}
	walk(node.Left, ch)
	if node.Left == nil && node.Right == nil {
		ch <- node.Val
	}
	walk(node.Right, ch)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		walk(root1, ch1)
	}()

	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		walk(root2, ch2)
	}()

	for n1 := range ch1 {
		n2, ok := <-ch2
		if n2 != n1 || !ok {
			return false
		}
	}

	_, ok := <-ch2
	if ok {
		return false
	}

	return true
}

type NilInt struct {
	value int
	null  bool
}

func (n *NilInt) Value() interface{} {
	if n.null {
		return nil
	}
	return n.value
}

func NewInt(x int) NilInt {
	return NilInt{x, false}
}

func NewNil() NilInt {
	return NilInt{0, true}
}
