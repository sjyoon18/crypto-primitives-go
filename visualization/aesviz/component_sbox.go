package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
	"math/bits"
)

// Function VisualizeSBoxByte visualizes the byte to byte substitution
// done by the corresponding S-box.
func PrintSBoxByte(b byte) {
	out := aes.SubstituteByte(b)

	printTitle("AES S-box Byte Visualization")

	fmt.Printf("Input : 0x%02x (%08b)\n", b, b)
	fmt.Printf("Output: 0x%02x (%08b)\n", out, out)
	fmt.Printf("Bits changed: %d\n", bits.OnesCount8(b^out))
}
