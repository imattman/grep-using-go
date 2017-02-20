// go build -o gogrep
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pattern> <file>\n", os.Args[0])
		os.Exit(2)
	}

	pat := os.Args[1]
	file := os.Args[2]

	cnt, err := printMatchingLines(os.Stdout, pat, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(2)
	}

	if cnt > 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}

func printMatchingLines(out io.Writer, pat string, file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	matchCnt := 0
	scan := bufio.NewScanner(bufio.NewReader(f))
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, pat) {
			fmt.Fprintln(out, line)
			matchCnt++
		}
	}
	if scan.Err() != nil {
		return matchCnt, scan.Err()
	}

	return matchCnt, nil
}
