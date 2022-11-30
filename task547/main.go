package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

/*
Three ways of taking input

 1. fmt.Scanln(&input)

 2. reader.ReadString()

 3. scanner.Scan()

Here we recommend using bufio.NewScanner
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter connected matrix: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			i, err := solve(text)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(i)
		} else {
			break
		}
	}
}

func solve(input string) (int, error) {
	isConnected, err := getSlice(input)
	if err != nil {
		return 0, err
	}

	n := len(isConnected)
	if n == 0 {
		return 0, nil
	}
	viewed := make([]int, n)

	var dfs func(node int)
	dfs = func(node int) {
		if viewed[node] == 0 {
			viewed[node] = 1
			for j := 0; j < n; j++ {
				if isConnected[node][j] == 1 {
					dfs(j)
				}
			}
		} else {
			return
		}
	}

	var count int
	for i := 0; i < n; i++ {
		if viewed[i] == 0 {
			dfs(i)
			count++
		}
	}

	return count, nil
}

// getSlice modify input string to slice of slices integers
// Example: [[1,1,0],[1,1,0],[0,0,1]]
func getSlice(input string) ([][]int, error) {
	res := make([][]int, 0)
	var prev rune

	if input == "" {
		return res, errors.New("empty input string")
	}

	tmp := make([]int, 0)

	for _, w := range input {
		if (w == '[' && prev == '[') || (w == ']' && prev == ']') {
			prev = w
			continue
		}
		if w == ',' || w == ' ' {
			prev = w
			continue
		}
		if w == '1' {
			tmp = append(tmp, 1)
			prev = w
			continue
		}
		if w == '0' {
			tmp = append(tmp, 0)
			prev = w
			continue
		}
		if w == '[' {
			tmp = make([]int, 0)
			prev = w
			continue
		}
		if w == ']' {
			res = append(res, tmp)
			prev = w
			continue
		}
		prev = w
	}

	if !validSlice(res) {
		return res, errors.New("wrong slice format. Example: [[1,1,0],[1,1,0],[0,0,1]]")
	}

	return res, nil
}

// validSlice check that input slice is correct
func validSlice(input [][]int) bool {
	n := len(input)
	for i := 0; i < n; i++ {
		if len(input[i]) != n {
			return false
		}

		for j := 0; j < n; j++ {
			if !(input[i][j] == 0 || input[i][j] == 1) {
				return false
			}
		}
	}

	return true
}
