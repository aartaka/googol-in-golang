package main

import "fmt"

func main() {
	var googol float64 = 1
	for i := 0; i < 100; i += 1 {
		googol *= 10
	}
	fmt.Printf("%e\n", googol)
}
