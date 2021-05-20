package main

/* Googol is a sequence of bytes encoding Googol.
 *
 * Primary takeaways:
 *
 * - Binary encoding of Googol is 333 bits long.
 * - switch is a convenient control structure, much like Lisps' cond.
 * - "% x" format specifier puts pretty spaces between bytes.
 * - There are several ways to make byte slices: new, make and direct
 *   initialization:
 *   - new returns pointer to uninitialized memory.
 *   - make fills the memory and returns the object itself.
 *   - direct initialization makes almost no sense with the two above,
 *     but it can be handy when you know the initial contents.
 */

import "fmt"

func main() {
	googol_bytes := []byte{0x12, 0x49, 0xAD, 0x25, 0x94, 0xC3, 0x7C, 0xEB,
		0x0B, 0x27, 0x84, 0xC4, 0xCE, 0x0B, 0xF3, 0x8A, 0xCE, 0x40, 0x8E, 0x21,
		0x1A, 0x7C, 0xAA, 0xB2, 0x43, 0x08, 0xA8, 0x2E, 0x8F, 0x10, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	fmt.Printf("% x\n", googol_bytes)
}
