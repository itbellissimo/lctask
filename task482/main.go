package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var (
		s = "5F3Z-2e-9-w"
		k = 4
	)

	fmt.Println(s, k)
	res := licenseKeyFormatting(s, k)
	fmt.Println(res)
}

func licenseKeyFormatting(s string, k int) string {
	var buf bytes.Buffer
	fstIsCorrect := false
	cnt := 0
	firstBlock := true

	for i, char := range s {
		// skip first incorrect symbols. If we don't need format wrong code, it must return error
		if fstIsCorrect == false && char != '-' {
			fstIsCorrect = true
		}

		if !fstIsCorrect {
			continue
		}

		if firstBlock && char == '-' && cnt <= k {
			buf.WriteByte(byte('-'))
			firstBlock = false
			cnt = 0
			continue
		}

		if char == '-' {
			continue
		}

		if char >= 'a' && char <= 'z' {
			char = char + 'Z' - 'z'
		}
		buf.WriteByte(byte(char))
		cnt++

		if cnt == k {
			if i != len(s)-1 {
				buf.WriteByte(byte('-'))
			}
			if firstBlock {
				firstBlock = false
			}
			cnt = 0
		}
	}

	b := buf.Bytes()
	if len(b) > 0 && b[len(b)-1] == '-' {
		b = b[:len(b)-1]
	}

	return string(b)
}

// licenseKeyFormattingUseStrings is realistion licenseKeyFormatting but use standard package strings
func licenseKeyFormattingUseStrings(s string, k int) string {
	s = strings.Trim(s, "-")
	s = strings.ToUpper(s)

	index := strings.Index(s, "-")
	prefix := ""
	if index != -1 && index <= k-1 {
		prefix = s[:index+1]
		s = s[index:]
	}

	s = strings.Replace(s, "-", "", -1)
	mod := k
	for mod < len(s) {
		s = s[:mod] + "-" + s[mod:]
		mod += k + 1
	}

	return prefix + s
}
