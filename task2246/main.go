package main

import (
	"fmt"
)

func main() {
	// parent, s := []int{-1, 0, 0, 1, 1, 2}, "abacbe"
	parent, s := []int{-1, 0, 0, 0}, "aabc"

	t := groupByParent(parent)
	fmt.Println(t)

	res := longestPath(parent, s)
	fmt.Println(res)
}

// groupByParent return indexes group by parent
func groupByParent(parent []int) map[int][]int {
	// 0 => {1,2}, 1 => {3,4}, 2 => {5}
	indexes := make(map[int][]int)
	for k, v := range parent {
		if v != -1 {
			indexes[v] = append(indexes[v], k)
		}
	}

	return indexes
}

func dfs(u int, indexes map[int][]int, s string, max *int) int {
	n1, n2 := 0, 0
	for _, v := range indexes[u] {
		maxDFS := dfs(v, indexes, s, max)
		if s[v] != s[u] {
			if maxDFS > n1 {
				n2 = n1
				n1 = maxDFS
			} else if maxDFS > n2 {
				n2 = maxDFS
			}
		}
	}
	if n1+n2+1 > *max {
		*max = n1 + n2 + 1
	}
	return n1 + 1
}

func longestPath(parent []int, s string) int {
	max := 0

	// 0 => {1,2}, 1 => {3,4}, 2 => {5}
	indexes := make(map[int][]int)
	for k, v := range parent {
		if v != -1 {
			indexes[v] = append(indexes[v], k)
		}
	}

	dfs(0, indexes, s, &max)
	return max
}
