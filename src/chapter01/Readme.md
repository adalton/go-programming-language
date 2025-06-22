1. Modify the `echo` program to also print `os.Args[0]`, the name of the command
   that invoked it.

   [Solution](./exercise01/exercise01.go)

   Sample output:

   ```
   $ go run exercise01.go this is a test
   /tmp/go-build601632563/b001/exe/exercise01 this is a test
   ```

2. Modify the `echo` program to print the index and value of each of its
   arguments, one per line.

   [Solution](./exercise03/exercise03.go)

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

   [Implemenetations](./exercise03/exercise03.go)
   [Benchmark program](./exercise03/exercise03_test.go)

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

   [Solution](./exercise04/exercise04.go)

5. Change the Lissajous program's color palette to green on black, for added
   authenticity.  To create the web color `#RRGGBB`, use
   `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, where each pair of hexadecimal digits
   represents the intensity of the reg, green, or blue component of the pixel.

   [Solution](./exercise05/exercise05.go)
   [Sample Output](./exercise05/sampleOutput.gif)

6. Modify the Lissajous progrm to produce images in multiple colors by adding
   more values to `palette` and then displaying them by changing the third
   argument to `SetColorIndex` in some interesting way.

   In this case, I updated the program to change the foreground color between
   red, green, and blue on each iteration.

   [Solution](./exercise06/exercise06.go)
   [Sample Output](./exercise06/sampleOutput.gif)

