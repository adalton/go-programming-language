// Copyright © 2016, 2019-2020 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.

package exercise03

import (
	"fmt"
	"os"
	"strings"
)

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
