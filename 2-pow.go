package main

/* Primary takeaways:
 *
 * - Go has math library.
 * - Pow is quite conventional -- math.Pow(base, power).
 * - Initializing declaration (:=, walrus operator) allows
 *   avoiding obvious types.
 */

import (
	"fmt"
	"math"
)

func main() {
	googol := math.Pow(10, 100)
	fmt.Printf("%e\n", googol)
}
