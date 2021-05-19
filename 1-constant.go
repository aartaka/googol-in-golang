package main

/* Googol is a float64 constant.
 *
 * Primary takeaways:
 *
 * - It's hard to provide Googol constant -- it's too big for anything
 *   except float64.
 * - Integer literal is not converted to float on oveflow, so we need to:
 *   - Use explicit typing,
 *   - Set .0 (or simply .) in the end, or
 *   - Use exponential notation.
 * - In any way, Googol is too big to be exact, thus the output:
 *   10000000000000000159028911097599180468360808563945281389781327557747838772170381060813469985856815104.000000
 */

import "fmt"

func main() {
	var googol_int float64 = 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	googol_float := 10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.0
	googol_exp := 1.000000e+100
	fmt.Printf("%f = %f = %e\n", googol_int, googol_float, googol_exp)
}
