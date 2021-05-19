package main

/* Primary takeaways:
 *
 * - Random numbers are deterministic in Go (this program always
 *   results in 1.647999e+100).
 * - The choice between exponential, normal, and logarithmic
 *   distribution is a hard one :D
 * - The compound version of conditional (`if init; cond { body}`)
 *   comes in handy, but only if initialization is short. Otherwise,
 *   semicolon gets lost and it puts more burden on your brain.
 * - Half (or more) of uniformly distributed random floats are below 1
 *   in their absolute value.
 */

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	googol := 1.0
	for googol < 1.000000e+100 {
		if tmp := math.Abs(rand.ExpFloat64()); tmp > 1 {
			googol *= tmp
		}
	}
	fmt.Printf("%e\n", googol)
}
