package main

/* Primary takeaways:
 *
 * - Go's fmt.Scanf is much like C's scanf.
 * - Scanning to float64 is possible with %f and %g (I only expected
 *   the latter).
 * - fmt.Println(something) = fmt.Printf("%v\n", something).
 * - I need to reflect on my fear of long lines. It's unproductive (is it?)
 */

import "fmt"

func main() {
	var googol float64
	fmt.Sscanf("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", "%f", &googol)
	fmt.Println(googol)
}
