package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLicenseKeyFormatting calls licenseKeyFormatting
func TestLicenseKeyFormatting(t *testing.T) {

	type licenseKeyFormattingTest struct {
		s      string
		k      int
		output string
	}

	var licenseKeyFormattingTests = []licenseKeyFormattingTest{
		{
			s:      "5F3Z-2e-9-w",
			k:      4,
			output: "5F3Z-2E9W",
		},
		{
			s:      "2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "-2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "----2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "2-5g-3-J",
			k:      3,
			output: "2-5G3-J",
		},
		{
			s:      "2gghnhh-5gmmmmm-3-Jkkkkkkk",
			k:      2,
			output: "2G-GH-NH-H5-GM-MM-MM-3J-KK-KK-KK-K",
		},
		{
			s:      "2g-5gmm-3k",
			k:      3,
			output: "2G-5GM-M3K",
		},
		{
			s:      "2g-5gmm-3k----",
			k:      3,
			output: "2G-5GM-M3K",
		},
		{
			s:      "",
			k:      1,
			output: "",
		},
	}

	for _, test := range licenseKeyFormattingTests {
		res := licenseKeyFormatting(test.s, test.k)
		assert.Equal(t, test.output, res)
	}
}

// TestLicenseKeyFormattingUseStrings calls licenseKeyFormattingUseStrings
func TestLicenseKeyFormattingUseStrings(t *testing.T) {

	type licenseKeyFormattingTest struct {
		s      string
		k      int
		output string
	}

	var licenseKeyFormattingTests = []licenseKeyFormattingTest{
		{
			s:      "5F3Z-2e-9-w",
			k:      4,
			output: "5F3Z-2E9W",
		},
		{
			s:      "2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "-2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "----2-5g-3-J",
			k:      2,
			output: "2-5G-3J",
		},
		{
			s:      "2-5g-3-J",
			k:      3,
			output: "2-5G3-J",
		},
		{
			s:      "2gghnhh-5gmmmmm-3-Jkkkkkkk",
			k:      2,
			output: "2G-GH-NH-H5-GM-MM-MM-3J-KK-KK-KK-K",
		},
		{
			s:      "2g-5gmm-3k",
			k:      3,
			output: "2G-5GM-M3K",
		},
		{
			s:      "2g-5gmm-3k----",
			k:      3,
			output: "2G-5GM-M3K",
		},
		{
			s:      "",
			k:      1,
			output: "",
		},
	}

	for _, test := range licenseKeyFormattingTests {
		res := licenseKeyFormattingUseStrings(test.s, test.k)
		assert.Equal(t, test.output, res)
	}
}
