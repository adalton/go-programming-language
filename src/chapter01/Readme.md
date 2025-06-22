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

5. Change the Lissajous program's color palette to green on black, for added
   authenticity.  To create the web color `#RRGGBB`, use
   `color.RGBA{0xRR, 0xGG, 0xBB, 0xff}`, where each pair of hexadecimal digits
   represents the intensity of the reg, green, or blue component of the pixel.

   ```go
   // Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
   // License: https://creativecommons.org/licenses/by-nc-sa/4.0/
   //
   // Modified by Andy Dalton to implement exercise solution.
   
   package main
   
   import (
   	"image"
   	"image/color"
   	"image/gif"
   	"io"
   	"math"
   	"math/rand"
   	"os"
   )
   
   var palette = []color.Color{color.Black, color.RGBA{0, 0xFF, 0, 0xFF}}
   
   const (
   	whiteIndex = 0 // first color in palette
   	blackIndex = 1 // next color in palette
   )
   
   func main() {
   	lissajous(os.Stdout)
   }
   
   func lissajous(out io.Writer) {
   	const (
   		cycles  = 5     // number of complete x oscillator revolutions
   		res     = 0.001 // angular resolution
   		size    = 100   // image canvas covers [-size..+size]
   		nframes = 64    // number of animation frames
   		delay   = 8     // delay between frames in 10ms units
   	)
   
   	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
   	anim := gif.GIF{LoopCount: nframes}
   	phase := 0.0 // phase difference
   
   	for i := 0; i < nframes; i++ {
   		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
   		img := image.NewPaletted(rect, palette)
   
   		for t := 0.0; t < cycles*2*math.Pi; t += res {
   			x := math.Sin(t)
   			y := math.Sin(t*freq + phase)
   
   			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
   		}
   
   		phase += 0.1
   		anim.Delay = append(anim.Delay, delay)
   		anim.Image = append(anim.Image, img)
   	}
   
   	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
   }
   ```

6. Modify the Lissajous progrm to produce images in multiple colors by adding
   more values to `palette` and then displaying them by changing the third
   argument to `SetColorIndex` in some interesting way.

   In this case, I updated the program to change the foreground color between
   red, green, and blue on each iteration.

   ```go
   // Copyright © 2016, 2019 Alan A. A. Donovan & Brian W. Kernighan.
   // License: https://creativecommons.org/licenses/by-nc-sa/4.0/
   //
   // Modified by Andy Dalton to implement exercise solution.
   
   package main
   
   import (
   	"image"
   	"image/color"
   	"image/gif"
   	"io"
   	"math"
   	"math/rand"
   	"os"
   )
   
   var palette = [][]color.Color{
   	{color.Black, color.RGBA{0xFF, 0, 0, 0xFF}},
   	{color.Black, color.RGBA{0, 0xFF, 0, 0xFF}},
   	{color.Black, color.RGBA{0, 0, 0xFF, 0xFF}},
   }
   
   const (
   	whiteIndex = 0 // first color in palette
   	blackIndex = 1 // next color in palette
   )
   
   func main() {
   	lissajous(os.Stdout)
   }
   
   func lissajous(out io.Writer) {
   	const (
   		cycles  = 5     // number of complete x oscillator revolutions
   		res     = 0.001 // angular resolution
   		size    = 100   // image canvas covers [-size..+size]
   		nframes = 64    // number of animation frames
   		delay   = 8     // delay between frames in 10ms units
   	)
   
   	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
   	anim := gif.GIF{LoopCount: nframes}
   	phase := 0.0 // phase difference
   
   	for i := 0; i < nframes; i++ {
   		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
   		img := image.NewPaletted(rect, palette[i%3])
   
   		for t := 0.0; t < cycles*2*math.Pi; t += res {
   			x := math.Sin(t)
   			y := math.Sin(t*freq + phase)
   
   			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
   		}
   
   		phase += 0.1
   		anim.Delay = append(anim.Delay, delay)
   		anim.Image = append(anim.Image, img)
   	}
   
   	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
   }
   ```

