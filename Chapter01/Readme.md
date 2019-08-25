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
   	// Here i is the index into the slice, not the array
   	for i, arg := range os.Args[1:] {
   		fmt.Printf("os.Args[%d] = %s\n", i, arg)
   		s += sep + arg
   		sep = " "
   	}
   	fmt.Println(s)
   }
   ```

   Sample output:

   ```
   $ go run exercise02.go this is a test
   os.Args[0] = this
   os.Args[1] = is
   os.Args[2] = a
   os.Args[3] = test
   this is a test
   ```

3. Experiment to measure the difference in running time between our potentially
   inefficient versions and the one that uses `strings.Join`. (Section 1.6
   illustrates part of the `time` package, and Section 11.4 shows how to write
   benchmark tests for systematic performance evaulation.)

   Benchmark program:
   ```go
   // Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
   // License: https://creativecommons.org/licenses/by-nc-sa/4.0/
   //
   // Modified by Andy Dalton to implement exercise solution.
   
   package main
   
   import (
   	"fmt"
   	"os"
   	"strings"
   	"testing"
   )
   
   func echo2() {
   	s, sep := "", ""
   	for _, arg := range os.Args[1:] {
   		s += sep + arg
   		sep = " "
   	}
   	fmt.Println(s)
   }
   
   func echo3() {
   	fmt.Println(strings.Join(os.Args[1:], " "))
   }
   
   func BenchmarkEcho2(b *testing.B) {
   	for i := 0; i < b.N; i++ {
   		echo2()
   	}
   }
   
   func BenchmarkEcho3(b *testing.B) {
   	for i := 0; i < b.N; i++ {
   		echo3()
   	}
   }
   ```

   Sample run:

   ```
   $ go test -bench=. -- $(for ((i = 0; i < 100; ++i)); do echo "arg${i}"; done) | grep -v arg
   goos: linux
   goarch: amd64
      50000	     24708 ns/op
     200000	      7977 ns/op
   PASS
   ok  	_/go-programming-language/Chapter01/exercise03	3.167s
   ```

4. Modify `dup2` to print the names of all files in which each duplicated line
   occurs.

   Solution:

   ```go
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
   ```
