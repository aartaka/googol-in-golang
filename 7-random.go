package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	googol := 1.0
	for googol < 1.000000e+100 {
		if tmp := math.Abs(rand.ExpFloat64()); tmp > 1 {
			googol *= tmp
		}
	}
	fmt.Printf("%e\n", googol)
}
