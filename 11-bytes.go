package main

/* Googol is a sequence of bytes encoding Googol.
 *
 * Primary takeaways:
 *
 * - Binary encoding of Googol is 333 bytes long.
 * - switch is a convenient control structure, much like Lisps' cond.
 * - `for idx, val := range array` is a huge improvement over C. Is it
 *    extensible for user types, however?
 * - "% x" format specifier puts pretty spaces between bytes.
 * - There are several ways to make byte slices: new, make and direct initialization:
 *   - new returns pointer to uninitialized memory.
 *   - make fills the memory and returns the object itself.
 *   - direct initialization makes almost no sense with the two above,
 *     but it's just variable declaration and a loop over it's bytes.
 */

import "fmt"

func main() {
	googol_string := "100100100100110101101001001011001010011000011011111001110101100001011001001111000010011000100110011100000101111110011100010101100111001000000100011100010000100011010011111001010101010110010010000110000100010101000001011101000111100010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	googol_bytes := make([]byte, 333, 333)
	for idx, val := range googol_string {
		switch {
		case val == '1':
			googol_bytes[idx] = 1
		case val == '0':
			googol_bytes[idx] = 0
		}
	}
	fmt.Printf("% x\n", googol_bytes)
}
