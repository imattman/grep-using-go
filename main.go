// go build -o gogrep
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fatal(2, "Usage: %s <pattern> <file>\n", os.Args[0])
	}

	pat := os.Args[1]
	file := os.Args[2]

	err := printMatchingLines(pat, file)
	if err != nil {
		fatal(2, err.Error())
	}
}

func fatal(exitVal int, msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
	os.Exit(exitVal)
}

func printMatchingLines(pat string, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scan := bufio.NewScanner(bufio.NewReader(f))
	for scan.Scan() {
		line := scan.Text()
		if strings.Contains(line, pat) {
			fmt.Println(line)
		}
	}
	if scan.Err() != nil {
		return scan.Err()
	}

	return nil
}
