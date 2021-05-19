package main

/* Googol is 1 recursively multiplied by 10 (until it's Googol).
 *
 * Primary takeaways:
 *
 * - Recursion can go as deep as 100 frames in Go.
 * - But stack is not endless. Here's the error message I got when
     replacing multiplication with addition:
 *   `runtime: goroutine stack exceeds 1000000000-byte limit`
 * - Thus, there's no tail-call optimization in Go. What a shame :(
 */

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
