package main

import "fmt"

func main() {
	var googol float64 = 1
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			googol *= 10
		}
	}
	fmt.Printf("%e\n", googol)
}
