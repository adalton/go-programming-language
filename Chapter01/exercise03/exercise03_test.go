// Copyright © 2016, 2019-2020 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Modified by Andy Dalton to implement exercise solution.

package exercise03_test

import (
	"testing"

	"exercise03"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise03.echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		exercise03.echo3()
	}
}
