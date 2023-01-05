package main

import (
	"fmt"
	"github.com/itbellissimo/treenode"
)

func main() {
	treeStr := treenode.NewTreeString("21,7,14,1,1,2,2,3,3")
	root, err := treeStr.GetTreeNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	res := findTilt(root)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *treenode.TreeNode) int {
	var (
		ret int
		f   func(node *treenode.TreeNode) int
	)

	f = func(node *treenode.TreeNode) int {
		if node == nil {
			return 0
		}

		l := f(node.Left)
		r := f(node.Right)
		ret += abs(l - r)
		return l + r + node.Val
	}

	f(root)

	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
