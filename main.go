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
		fmt.Fprintf(os.Stderr, "Usage: %s <pattern> <file>\n", os.Args[0])
		// TODO: figure out why stderr isn't flushing with fmt.Errorf()
		// fmt.Errorf("Usage: %s <pattern> <file>\n", os.Args[0])
		// os.Stderr.Sync()
		os.Exit(2)
	}

	pat := os.Args[1]
	file := os.Args[2]

	err := printMatchingLines(pat, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
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
