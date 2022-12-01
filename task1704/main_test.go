package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

// TestHalvesAreAlike calls halvesAreAlike
func TestHalvesAreAlike(t *testing.T) {

	type halvesAreAlikeTest struct {
		s     string
		l     int
		strA  string
		strB  string
		alike bool
	}

	var halvesAreAlikes = []halvesAreAlikeTest{
		{
			s:     "aeoIoOuUIoOaeouU",
			l:     16,
			alike: true,
		},
		{
			s:     "book",
			l:     4,
			alike: true,
		},
		{
			s:     "books",
			l:     5,
			alike: false,
		},
		{
			s:     "textbook",
			l:     8,
			alike: false,
		},
		{
			s:     "bommbook",
			l:     8,
			alike: false,
		},

		{
			s:     "boOk",
			l:     4,
			alike: true,
		},
	}

	for _, test := range halvesAreAlikes {
		alike := halvesAreAlike(test.s)
		assert.Equal(t, test.l, utf8.RuneCountInString(test.s))
		assert.Equal(t, test.alike, alike)
	}
}
