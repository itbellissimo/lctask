package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		pattern = "abba"
		s       = "dog cat cat dog"
	)

	res := wordPattern(pattern, s)
	fmt.Println(res)
}

func wordPattern(pattern string, s string) bool {
	// string separate by whitespace to slice of words
	sl := strings.Split(s, " ")
	conformity := make(map[string]string)
	words := make(map[string]string)
	if len(pattern) != len(sl) {
		return false
	}

	// check template
	for i, p := range pattern {
		if w, ok := conformity[string(p)]; ok && w != sl[i] {
			return false
		}
		if existsP, ok := words[sl[i]]; ok && string(p) != existsP {
			return false
		}

		conformity[string(p)] = sl[i]
		words[sl[i]] = string(p)
	}

	if len(sl) < len(conformity) {
		return false
	}

	return true
}
