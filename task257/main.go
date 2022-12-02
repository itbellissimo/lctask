package main

import (
	"fmt"
)

func main() {
	// root = [1,2,3,null,5]
	data := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	res := binaryTreePaths(&data)
	fmt.Println(res)
}

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var path string

	rootPrint(root, &res, path)

	return res
}

func rootPrint(root *TreeNode, res *[]string, path string) {
	if root == nil {
		return
	}

	sep := ""
	if path != "" {
		sep = "->"
	}

	path += fmt.Sprintf("%s%d", sep, root.Val)

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}

	rootPrint(root.Left, res, path)
	rootPrint(root.Right, res, path)
}
