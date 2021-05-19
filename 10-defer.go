package main

/* Primary takeaways:
 *
 * - defer is fun, but not obvious.
 * - defer is ran only after return statement, so it I did
 *   fmt.Println(googol(&g)), I'd get 1. And I did.
 * - defer only runs functions, you cannot pass an arbitrary statement
 *   to it. It's a design desision, I guess.
 * - Pointers are much similar to C ones.
 * - But you cannot do *= on dereferenced pointer somewhy.
 */

import "fmt"

func multBy10 (num *float64) {
	*num = *num * 10
}

func googol (num *float64) float64 {
	for i := 0; i < 100; i++ {
		defer multBy10(num)
	}
	return *num
}

func main() {
	g := 1.0
	googol(&g)
	fmt.Println(g)
}
