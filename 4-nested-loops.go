package main

import "fmt"

func main() {
	googol := 1.0
	for i := 0; i < 10; i++ {
		for i := 0; i < 10; i++ {
			googol *= 10
		}
	}
	fmt.Printf("%e\n", googol)
}
