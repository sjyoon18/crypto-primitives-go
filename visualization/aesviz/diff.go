package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function PrintStateDifference displays the diffusion information
// between two AES states.
func PrintStateDifference(diff aes.StateDifference) {
	fmt.Printf(
		"Bit differences: %d\n", diff.BitDifferences,
	)

	fmt.Printf(
		"Byte differences: %d\n", diff.ByteDifferences,
	)

	fmt.Println("\nChanged bytes:")
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if diff.ChangedBytes[r][c] {
				fmt.Printf(" * ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println()
	}

	fmt.Println("\nBit changed per byte:")
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			fmt.Printf(" %d ", diff.ChangedBits[r][c])
		}
		fmt.Println()
	}
}

// Function PrintRoundDifference displays diffusion
// growth throughout an AES round.
func PrintRoundDifference(rdiff aes.RoundDifference) {
	fmt.Printf("\n====== Round %d ======\n", rdiff.Round)

	fmt.Println("\nStart State:")
	PrintStateDifference(rdiff.StartState)

	fmt.Println("\nAfter Subbytes:")
	PrintStateDifference(rdiff.AfterSubBytes)

	fmt.Println("\nAfter ShiftRows:")
	PrintStateDifference(rdiff.AfterShiftRows)

	if rdiff.Round != 10 {
		fmt.Println("\nAfter MixColumns:")
		PrintStateDifference(rdiff.AfterMixColumns)
	}

	fmt.Println("\nAfter AddRoundKey:")
	PrintStateDifference(rdiff.AfterAddRoundKey)
}

func PrintTraceDifference(diff aes.EncryptionTraceDifference) {
	fmt.Println("AES Differential Trace")
	fmt.Println("========================")

	fmt.Println("\nPlaintext Difference:")
	PrintStateDifference(diff.PlaintextDifference)

	fmt.Println("\nAfter Initial Key Addition:")
	PrintStateDifference(diff.InitialAddRoundKeyDifference)

	for _, round := range diff.Rounds {
		PrintRoundDifference(round)
	}

	fmt.Println("\nCiphertext Difference:")
	PrintStateDifference(diff.CiphertextDifference)
}
