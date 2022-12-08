package main

import (
	"fmt"
	"math"
)

func main() {
	mat := [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}

	fmt.Println(mat)
	mat2 := updateMatrix(mat)
	fmt.Println(mat2)

	mat3 := updateMatrixFromSite(mat)
	fmt.Println(mat3)
}

type point struct {
	x int
	y int
}

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
	}

	// find zero points coordinates.
	zero := make([]*point, 0, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				r[i][j] = 0
				zero = append(zero, &point{
					x: i,
					y: j,
				})
			} else {
				// init to max possible delta.
				// 1 <= m, n <= 10^4
				r[i][j] = math.MaxInt //10000
			}
		}
	}

	if len(zero) == 0 {
		return mat
	}

	// for each point find minimum delta coordinates zero points and current point
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, v := range zero {
				t := int(math.Abs(float64(i-v.x)) + math.Abs(float64(j-v.y)))
				if r[i][j] > t {
					r[i][j] = t
				}
			}
		}
	}

	return r
}

var directions = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// updateMatrixFromSite the solution is from comments leetcode. Not mine.
func updateMatrixFromSite(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				mat[i][j] = math.MaxInt64
			}
		}
	}
	//fmt.Println("queue: ", queue)
	//fmt.Println("mat: ", mat)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			x, y := cur[0]+dir[0], cur[1]+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n || mat[x][y] <= mat[cur[0]][cur[1]] {
				continue
			}
			queue = append(queue, [2]int{x, y})
			mat[x][y] = mat[cur[0]][cur[1]] + 1
		}
	}

	return mat
}
