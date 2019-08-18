// Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// Here i is the index into the slice, not the array
	for i, arg := range os.Args[1:] {
		fmt.Printf("os.Args[%d] = %s\n", i, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
