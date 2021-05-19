package main

/* Primary takeaways:
 *
 * - for loop with one argument is while loop. Is it convenient?
 * - Iteration seems to be the shortest way of problem-solving in Go.
 */

import "fmt"

func main() {
	googol := 1.0
	for googol < 1.000000e+100 {
		googol *= 10
	}
	fmt.Printf("%e\n", googol)
}
