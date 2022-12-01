package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s = "book"

	res := halvesAreAlike(s)
	fmt.Println(s, res)
}

func halvesAreAlike(s string) bool {
	l := utf8.RuneCountInString(s)

	if l%2 != 0 {
		return false
	}
	r := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
	var (
		countA int
		countB int
	)
	for i, v := range s {

		if _, ok := r[v]; i < l/2 && ok {
			countA++
		} else if ok {
			countB++
		}
	}

	return countA == countB
}
