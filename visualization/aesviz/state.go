package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function PrintState prints the AES state input as a 4x4 matrix.
func PrintState(state aes.State) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Printf("%02x ", state[r][c])
		}
		fmt.Println()
	}
	fmt.Println()
}

// Function PrintStateColumns prints the columns of the AES state.
func PrintStateColumns(state aes.State) {
	for c := 0; c < 4; c++ {
		fmt.Printf("column %d ", c)
		for r := 0; r < 4; r++ {
			fmt.Printf("%02x ", state[r][c])
		}
		fmt.Println()
	}
	fmt.Println()
}
