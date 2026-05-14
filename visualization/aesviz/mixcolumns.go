package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function PrintMixColumn prints the input and output
// of AES MixColumns for a single 4-byte AES column.
func PrintMixColumn(column [4]byte) {
	mixed := aes.MixSingleColumn(column)

	printTitle("AES MixColumns Visualization:")

	fmt.Println("Input column:")
	for i := 0; i < 4; i++ {
		fmt.Printf("  [%d] %02x\n", i, column[i])
	}

	fmt.Println("\nOutput column:")
	for i := 0; i < 4; i++ {
		fmt.Printf("  [%d] %02x\n", i, mixed[i])
	}
}
