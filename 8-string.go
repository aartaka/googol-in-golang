package main

/* Primary takeaways:
 *
 * - String are immutable.
 * - String concatenation is simply a + between two strings.
 * - Concatenation creates new strings and thus is slow and memory-inefficient.
 * - To make string mutable and make it fast, you need to use either
 *   string.Builder or bytes.Buffer.
 * - This outputs a perfect Googol with 100 zeros. Probably the only
 *   true program in this series xD
 */

import (
	"fmt"
	"bytes"
)

func main() {
	googol := bytes.NewBufferString("1")
	for i := 0; i < 100; i++  {
		googol.WriteString("0")
	}
	fmt.Printf("%s\n", googol)
}
