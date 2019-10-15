// Exercise 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dup struct {
	texts map[string]int
	path  string
}

func main() {
	counts := make(map[Dup]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[Dup]int, path string) {
	input := bufio.NewScanner(f)
	var dup Dup
	dup.path = path
	for input.Scan() {
		dup.text = input.Text()
		counts[dup]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
