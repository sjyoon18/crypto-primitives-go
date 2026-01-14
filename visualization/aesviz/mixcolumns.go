package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function VisualizeMixColumn visualizes the transformation of an AES column
// under MixColumns.
// The input column is placed into column 0 of the AES state while
// the other columns are set to zero to isolate MixColumns behavior.
func VisualizeMixColumn(col [4]byte) {
	fmt.Printf("Input column:\n")
	for i := 0; i < 4; i++ {
		fmt.Printf("  [%d] %02x\n", i, col[i])
	}

	state := aes.State{
		{col[0], 0, 0, 0},
		{col[1], 0, 0, 0},
		{col[2], 0, 0, 0},
		{col[3], 0, 0, 0},
	}

	aes.DebugMixColumns(&state)

	fmt.Println("\nOutput column:")
	for i := 0; i < 4; i++ {
		fmt.Printf("  [%d] %02x\n", i, state[i][0])
	}
}

// Function VisualizeMixColumnInfluence visualizes the effect of changing a single byte in an AES column
// in MixColumns by implementing VisualizeMixColumn.
func VisualizeMixColumnInfluence(bytePos int, value byte) {
	if bytePos < 0 || bytePos > 3 {
		panic("bytePos must be in range [0,3]")
	}
	var col [4]byte
	col[bytePos] = value
	VisualizeMixColumn(col)
}
