package main

import "fmt"

func googol (number float64) float64 {
	if number >= 1.000000e+100 {
		return number
	}
	return googol(number * 10)
}

func main() {
	fmt.Printf("%e\n", googol(1))
}
