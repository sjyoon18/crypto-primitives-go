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
	printBoolMatrix(
		diff.ChangedBytes,
		" * ",
		" . ",
	)

	fmt.Println("\nBit changed per byte:")
	printUint8Matrix(
		diff.ChangedBits,
		" %d ",
	)
}

// Function PrintRoundDifference displays diffusion
// growth throughout an AES round.
func PrintRoundDifference(rdiff aes.RoundDifference) {
	printRoundHeader(rdiff.Round)

	printSection("Start State:")
	PrintStateDifference(rdiff.StartState)

	printSection("After SubBytes:")
	PrintStateDifference(rdiff.AfterSubBytes)

	printSection("After ShiftRows:")
	PrintStateDifference(rdiff.AfterShiftRows)

	if rdiff.Round != 10 {
		printSection("After MixColumns:")
		PrintStateDifference(rdiff.AfterMixColumns)
	}

	printSection("After AddRoundKey:")
	PrintStateDifference(rdiff.AfterAddRoundKey)
}

func PrintTraceDifference(diff aes.EncryptionTraceDifference) {
	printTitle("AES Differential Trace")

	printSection("Plaintext Difference:")
	PrintStateDifference(diff.PlaintextDifference)

	printSection("After Initial AddRoundKey:")
	PrintStateDifference(diff.InitialAddRoundKeyDifference)

	for _, round := range diff.Rounds {
		PrintRoundDifference(round)
	}

	printSection("Ciphertext Difference:")
	PrintStateDifference(diff.CiphertextDifference)
}
