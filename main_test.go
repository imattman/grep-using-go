package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintMatchingLines(t *testing.T) {
	tests := []struct {
		pattern  string
		file     string
		matchYes []string // matching lines expected to contain these strings
		matchNo  []string // matching lines are not expected to contain these
	}{
		{
			pattern:  "fish",
			file:     "testdata/fish.txt",
			matchYes: []string{"One", "two", "red", "blue"},
			matchNo:  []string{"ONE", "TWO"},
		},
		{
			pattern:  "not-in-text",
			file:     "testdata/fish.txt",
			matchYes: []string{},
			matchNo:  []string{"One", "two", "red", "blue", "ONE", "TWO", "RED", "BLUE"},
		},
		{
			pattern:  " ", // space expected in all lines with text
			file:     "testdata/fish.txt",
			matchYes: []string{"One", "two", "red", "blue", "ONE", "TWO", "RED", "BLUE"},
			matchNo:  []string{},
		},
		{
			pattern:  " ",
			file:     "testdata/empty.txt",
			matchYes: []string{},
			matchNo:  []string{},
		},
	}

	for _, test := range tests {
		var buf bytes.Buffer
		printMatchingLines(&buf, test.pattern, test.file)
		capOut := buf.String()

		for _, exp := range test.matchYes {
			if !strings.Contains(capOut, exp) {
				t.Errorf("%q in %s: expected match of line with %q\n\n%s",
					test.pattern, test.file, exp, capOut)
			}
		}

		for _, notExp := range test.matchNo {
			if strings.Contains(capOut, notExp) {
				t.Errorf("%q in %s: expected no match of line with %q\n\n%s",
					test.pattern, test.file, notExp, capOut)
			}
		}
	}
}
