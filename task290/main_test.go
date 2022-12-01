package main

import (
	"testing"
)

// TestWordPattern calls wordPattern
// for a valid return value.
func TestWordPattern(t *testing.T) {

	type wordPatternTest struct {
		pattern string
		s       string
		result  bool
	}

	var wordPatternTests = []wordPatternTest{
		{
			pattern: "abba",
			s:       "dog cat cat dog",
			result:  true,
		},
		{
			pattern: "abba",
			s:       "dog cat cat fish",
			result:  false,
		},
		{
			pattern: "aaaa",
			s:       "dog cat cat dog",
			result:  false,
		},
		{
			pattern: "abban",
			s:       "dog cat cat dog dog",
			result:  false,
		},
		{
			pattern: "abbaa",
			s:       "dog cat cat dog dog",
			result:  true,
		},
		{
			pattern: "ab",
			s:       "dog cat cat dog dog",
			result:  false,
		},
		{
			pattern: "abba",
			s:       "dog cat",
			result:  false,
		},
	}

	for _, test := range wordPatternTests {
		if output := wordPattern(test.pattern, test.s); output != test.result {
			t.Errorf("Pattern: %s, String: %s. Output %t not equal to expected %t",
				test.pattern, test.s, output, test.result)
		}
	}
}
