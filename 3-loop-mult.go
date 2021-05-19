package main

/* Googol is 10 multiplied by itself 100 times.
 *
 * Primary takeaways:
 *
 * - Go's loop is quite the same as in C, but absence of parentheses
 *   makes it look less structured. Easy to get lost.
 * - Increment operator (++) work in loops (I heard it's quite restricted).
 * - You always need to use brackets for loop body. Meh...
 * - If you initialize the number with a float literal, then it will
 *   be float64 (or at least it was the case in this snippet).
 */

import "fmt"

func main() {
	googol := 1.0
	for i := 0; i < 100; i++ {
		googol *= 10
	}
	fmt.Printf("%e\n", googol)
}
