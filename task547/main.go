package main

import (
	"bufio"
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
		fmt.Print("Enter Text: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()

		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			fmt.Println(solve(text))
		} else {
			break
		}
	}
}

func solve(input string) string {
	result := input
	return result
}
