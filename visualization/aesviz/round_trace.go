package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function TraceEncryption traces and prints the AES-128 encryption process
// of the 16 byte input (AES state).
func TraceEncryption(plaintext [16]byte, key [16]byte) {
	state := aes.BytesToState(plaintext)
	roundKeys := aes.DebugKeySchedule(key)

	fmt.Println("AES-128 Encryption Trace")
	fmt.Println("==========================")
	fmt.Println("Initial state:")
	PrintState(state)

	// Round 0
	aes.AddRoundKey(&state, roundKeys[:], 0)
	fmt.Println("\nRound 0:")
	PrintState(state)

	// Round 1 - 9
	for round := 1; round <= 9; round++ {
		fmt.Printf("\nRound %d:", round)

		aes.DebugSubBytes(&state)
		fmt.Println("After SubBytes:")
		PrintState(state)

		aes.DebugShiftRows(&state)
		fmt.Println("After ShiftRows:")
		PrintState(state)

		aes.DebugMixColumns(&state)
		fmt.Println("After MixColumns:")
		PrintState(state)

		aes.AddRoundKey(&state, roundKeys[:], round)
		fmt.Println("After AddRoundKey:")
		PrintState(state)
	}

	//Round 10
	fmt.Println("\nRound 10:")

	aes.DebugSubBytes(&state)
	fmt.Println("After SubBytes:")
	PrintState(state)

	aes.DebugShiftRows(&state)
	fmt.Println("After ShiftRows:")
	PrintState(state)

	aes.AddRoundKey(&state, roundKeys[:], 10)
	fmt.Println("After AddRoundKey (Final):")
	PrintState(state)
}
