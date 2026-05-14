package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function PrintState prints the AES state input as a 4x4 matrix.
func PrintState(state aes.State) {
	printByteMatrix(state, "%02x ")
	fmt.Println()
}
