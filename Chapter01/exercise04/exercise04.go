// Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.


// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure that maintains a count of the number of times a line has been
// seen and a slice containing the filenames in which they were seen.
type lineInfo struct {
	count int
	fileNames []string
}

func main() {
	counts := make(map[string]*lineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\n", n.count, line)
			fmt.Printf("\t%s\n", strings.Join(n.fileNames, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]*lineInfo) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = &lineInfo{ }
		}

		counts[input.Text()].count++
		counts[input.Text()].fileNames = append(counts[input.Text()].fileNames,
		                                        f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
