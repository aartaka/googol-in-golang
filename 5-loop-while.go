package main

import "fmt"

func main() {
	googol := 1.0
	for googol < 1.000000e+100 {
		googol *= 10
	}
	fmt.Printf("%e\n", googol)
}
