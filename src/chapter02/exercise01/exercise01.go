package main

import (
	"fmt"
	"exercise2_1/tempconv"
)

func main() {
	c := tempconv.AbsoluteZeroC

	fmt.Printf("%v = %v = %v\n", c, tempconv.CToF(c), tempconv.CToK(c))
}
