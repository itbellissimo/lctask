package main

import (
	"fmt"
)

func main() {
	grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
	grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}

	fmt.Println(grid1, grid2)
	res := countSubIslands(grid1, grid2)
	fmt.Println(res)
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var count int

	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if grid1[i][j] == 1 && grid2[i][j] == 1 {
				isMatching := true
				dfs(&grid1, &grid2, i, j, &isMatching)
				//fmt.Println("-----------------------------------------------------")
				//fmt.Println(grid2)
				//fmt.Println(i, j, isMatching)
				//fmt.Println("-----------------------------------------------------")
				if isMatching {
					count++
				}
			}
		}
	}
	return count
}

func dfs(g1 *[][]int, g2 *[][]int, i int, j int, isMatching *bool) {
	grid1, grid2 := *g1, *g2
	if i < 0 || i >= len(grid2) || j < 0 || j >= len(grid2[0]) || grid2[i][j] == 0 || grid2[i][j] == -1 {
		return
	}

	if grid1[i][j] == 0 && grid2[i][j] == 1 {
		*isMatching = false
	}
	grid2[i][j] = -1
	//fmt.Println(grid2)
	dfs(g1, g2, i+1, j, isMatching)
	dfs(g1, g2, i-1, j, isMatching)
	dfs(g1, g2, i, j+1, isMatching)
	dfs(g1, g2, i, j-1, isMatching)

}
