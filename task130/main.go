package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}

	fmt.Println(board)
	solve(board)
	fmt.Println(board)
}

func solve(board [][]byte) {
	if len(board) <= 2 && len(board[0]) <= 2 {
		return
	}

	i := 0
	for j := 0; j < len(board[0]); j++ {
		dfs(i, j, board)
	}

	i = len(board) - 1
	for j := 0; j < len(board[0]); j++ {
		dfs(i, j, board)
	}

	j := 0
	for i := 0; i < len(board); i++ {
		dfs(i, j, board)
	}

	j = len(board[0]) - 1
	for i := 0; i < len(board); i++ {
		dfs(i, j, board)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(i int, j int, board [][]byte) {
	if i < len(board) && j < len(board[0]) && board[i][j] == 'O' {
		board[i][j] = 'R'
	} else {
		return
	}

	if i+1 < len(board) {
		dfs(i+1, j, board)
	}
	if i-1 >= 0 {
		dfs(i-1, j, board)
	}
	if j+1 < len(board[0]) {
		dfs(i, j+1, board)
	}
	if j-1 >= 0 {
		dfs(i, j-1, board)
	}
}
