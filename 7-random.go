package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	googol := 1.0
	for googol < 1.000000e+100 {
		multiplier := math.Abs(rand.ExpFloat64())
		if multiplier > 1 {
			googol *= multiplier
		}
	}
	fmt.Printf("%e\n", googol)
}
