package main

import "fmt"

func main() {
	var googol float64 = 1
	for googol < 1.000000e+100 {
		googol *= 10
	}
	fmt.Printf("%e\n", googol)
}
