package main

/* Googol is 10 multiplied by itself 10x10 times.
 *
 * Primary takeaways:
 *
 * - Loops can be nested.
 * - You can use the same variable name in nested loops, and compiler
 *   with distinguish between two variables and rely on local scope.
 */

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
