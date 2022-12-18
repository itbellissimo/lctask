package main

import (
	"fmt"
	"github.com/itbellissimo/treenode"
)

func main() {
	treeStr := treenode.NewTreeString("8,3,10,1,6,null,14,null,null,4,7,13")
	node, err := treeStr.GetTreeNode()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(maxAncestorDiff(node))

	treeStr = treenode.NewTreeString("1,null,2,null,0,3")
	node, err = treeStr.GetTreeNode()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(maxAncestorDiff(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *treenode.TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *treenode.TreeNode, min, max int) int {
	if root == nil {
		return max - min
	}

	max = Max(max, root.Val)
	min = Min(min, root.Val)
	return Max(dfs(root.Left, min, max), dfs(root.Right, min, max))
}

func Max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
