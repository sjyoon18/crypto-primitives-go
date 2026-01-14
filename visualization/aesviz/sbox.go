package aesviz

import (
	"fmt"
	"math/bits"
)

// Function VisualizeSBoxByte visualizes the byte to byte substitution
// done by the corresponding S-box.
func VisualizeSBoxByte(b byte, sbox [256]byte) {
	out := sbox[b]
	fmt.Printf("Input : 0x%02x (%08b)\n", b, b)
	fmt.Printf("Output: 0x%02x (%08b)\n", out, out)
	fmt.Printf("Bits changed: %d\n", bits.OnesCount8(b^out))
}
