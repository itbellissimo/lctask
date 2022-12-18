package main

import (
	"fmt"
	"github.com/itbellissimo/treenode"
)

func main() {
	treeStr1 := treenode.NewTreeString("3,5,1,6,2,9,8,null,null,7,4")
	root1, err := treeStr1.GetTreeNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	treeStr2 := treenode.NewTreeString("3,5,1,6,2,9,8,null,null,7,4")
	root2, err := treeStr2.GetTreeNode()
	if err != nil {
		fmt.Println(err)
		return
	}

	mat2 := leafSimilar(root1, root2)
	fmt.Println(mat2)
}

func walk(node *treenode.TreeNode, ch chan int) {
	if node == nil {
		return
	}
	walk(node.Left, ch)
	if node.Left == nil && node.Right == nil {
		ch <- node.Val
	}
	walk(node.Right, ch)
}

func leafSimilar(root1 *treenode.TreeNode, root2 *treenode.TreeNode) bool {
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
