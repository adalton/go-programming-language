// Copyright © 2016, 2019-2020 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.
//
// go test -bench=. | grep -v timeout
// goos: linux
// goarch: amd64
// pkg: chapter01/exercise03
//   259570	      5173 ns/op
//   367957	      4664 ns/op
// PASS
// ok  	chapter01/exercise03	3.143s

package exercise03_test

import (
	"testing"

	"chapter01/exercise03"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise03.Echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise03.Echo3()
	}
}
