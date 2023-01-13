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

// longestPath
// Step 1. Find two longest path out of n children for each node.
// Step 2. Update global max value if longestPath1+longestPath2> max.
// Step 3. Return longest path for each child node to parent node using DFS.
//
// If adjacent node letters are not equal set pathlength for children = pathlength+1.
// If adjacent node letters are equal set pathlength for children = 0.
func longestPath(parent []int, s string) int {
	max := 0

	// 0 => {1,2}, 1 => {3,4}, 2 => {5}
	byParent := groupByParent(parent)

	// s "abacbe"
	var dfs func(int) int
	dfs = func(i int) int {
		lp1, lp2 := 0, 0 // lp1, lp2 - longest path 1 and 2
		for _, idx := range byParent[i] {
			dv := dfs(idx)
			if s[idx] != s[i] {
				if dv > lp1 {
					lp2 = lp1
					lp1 = dv
				} else if dv > lp2 {
					lp2 = dv
				}
			}
		}

		if lp1+lp2+1 > max {
			max = lp1 + lp2 + 1
		}

		return lp1 + 1
	}

	dfs(0)
	return max
}
