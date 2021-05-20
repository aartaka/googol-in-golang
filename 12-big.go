package main

/* Googol is a big.Int having an exact value of 10^100.
 *
 * Primary takeaways:
 *
 * - Arbitrary precision integers/floats/rationals (!) are
 *   built-in. The library is math/big.
 * - Int initialization happens with int64, so we cannot put Googol
 *   there right away.
 * - All the big operations are functions (methods, actually), so no
 *   operator overloading there.
 * - These methods modify the object they are called on, so no
 *   assignments these.
 * - Copying is only possible on big.Float (why?)
 * - There's native printing of big numbers with fmt! :D
 */

import (
	"fmt"
	"math/big"
)

func main() {
	googol := big.NewInt(1)
	ten := big.NewInt(10)
	for i := 0; i < 100; i++ {
		googol.Mul(googol, ten)
	}
	fmt.Println(googol)
}
