1. Modify the `echo` program to also print `os.Args[0]`, the name of the command
   that invoked it.

   ```go
   // Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
   // License: https://creativecommons.org/licenses/by-nc-sa/4.0/
   //
   // Modified by Andy Dalton to implement exercise solution.
   package main
   
   import (
   	"fmt"
   	"os"
   )
   
   func main() {
   	s, sep := "", ""
   	for _, arg := range os.Args[0:] {
   		s += sep + arg
   		sep = " "
   	}
   	fmt.Println(s)
   }
   ```

   Sample output:

   ```
   $ go run exercise01.go this is a test
   /tmp/go-build601632563/b001/exe/exercise01 this is a test
   ```

2. Modify the `echo` program to print the index and value of each of its
   arguments, one per line.

3. Experiment to measure the difference in running time between our potentially
   inefficient versions and the one that uses `strings.Join`. (Section 1.6
   illustrates part of the `time` package, and Section 11.4 shows how to write
   benchmark tests for systematic performance evaulation.)
