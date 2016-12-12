package main

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
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
		capOut, capErr := captureStdoutStderr(func() {
			printMatchingLines(test.pattern, test.file)
		}, t)

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

		if capErr != "" {
			t.Errorf("expected nothing in STDERR:\n%q", capErr)
		}
	}

}

// takes a zero-argument function and invokes the function while
// capturing all data written to stdout/error.
func captureStdoutStderr(f func(), t *testing.T) (string, string) {
	oldOut := os.Stdout
	oldErr := os.Stderr
	defer func() {
		os.Stdout = oldOut
		os.Stderr = oldErr
	}()

	var outBuf []byte
	var errBuf []byte
	var wg sync.WaitGroup
	outPipe := pipeForBuffer(&outBuf, &wg, t)
	errPipe := pipeForBuffer(&errBuf, &wg, t)

	os.Stdout = outPipe
	os.Stderr = errPipe
	f()

	outPipe.Close()
	errPipe.Close()
	wg.Wait() // allow the capture routines to finish

	return string(outBuf), string(errBuf)
}

func pipeForBuffer(buf *[]byte, wg *sync.WaitGroup, t *testing.T) *os.File {
	// Go doesn't use io.Writer for Stdout/Stderr
	// so we have to jump through some hoops using an
	// in-memory pipe to accomplish the capture
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("error with os.Pipe %q", err)
	}

	wg.Add(1)
	go func() {
		*buf, err = ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("error with os.Pipe %q", err)
		}
		wg.Done()
	}()

	return w
}
