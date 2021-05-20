package main

/* Googol is a number satisfying Googol interface.
 *
 * Primary takeaways:
 *
 * - You cannot define methods on the non-local types (like big.Int
 *   and float64). Ways out:
 *   - Define an alias.
 *   - Add a file with the necessary package (works only if you have
 *     all the build machinery set up).
 * - Nominal types are weird, but I can see a use in them.
 * - Floats are not really suitable for ==, as always, thus old trick
 *   of math.Abs and epsilons.
 * - The epsilon (difference between two neighboring floats) for
 *   Googol is freaking 1e+85!!!!
 * - Interfaces are cool, but I miss generics :_(
 */

import (
	"fmt"
	"math"
)

type float float64

type Googol interface {
	IsGoogol() bool
}

func (f float) IsGoogol() bool {
	return math.Abs(float64(f) - 1e+100) < 1e+85
}

func main() {
	var googol float = 1.0
	for !googol.IsGoogol() {
		googol *= 5
	}
	fmt.Println(googol)
}
