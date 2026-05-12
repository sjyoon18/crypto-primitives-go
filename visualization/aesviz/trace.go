package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

func PrintTrace(trace aes.EncryptionTrace) {
	fmt.Println("AES-128 Encryption")
	fmt.Println("====================")

	fmt.Println("\nPlaintext State:")
	PrintState(trace.Plaintext)

	fmt.Println("After Initial Key Addition:")
	PrintState(trace.InitialAddRoundKey)

	for _, round := range trace.Rounds {
		fmt.Printf("\n========== ROUND %d ==========\n", round.Round)

		fmt.Printf("\nStart State:")
		PrintState(round.StartState)

		fmt.Printf("Round Key:")
		PrintState(round.RoundKey)

		fmt.Printf("After SubBytes:")
		PrintState(round.AfterSubBytes)

		fmt.Printf("After ShiftRows:")
		PrintState(round.AfterShiftRows)

		if round.Round != 10 {
			fmt.Printf("After MixColumns:")
			PrintState(round.AfterMixColumns)
		}
		fmt.Printf("After AddRoundKey:")
		PrintState(round.AfterAddRoundKey)
	}

	fmt.Printf("\nCiphertext State:")
	PrintState(trace.Ciphertext)
}

/*
trace := aes.EncryptWithTrace(
	plaintext,
	key,
)

aesviz.PrintTrace(trace)
*/

func PrintRoundTrace(round aes.RoundTrace) {
	fmt.Printf("\n========== ROUND %d ==========\n", round.Round)

	fmt.Printf("\nStart State:")
	PrintState(round.StartState)

	fmt.Printf("Round Key:")
	PrintState(round.RoundKey)

	fmt.Printf("After SubBytes:")
	PrintState(round.AfterSubBytes)

	fmt.Printf("After ShiftRows:")
	PrintState(round.AfterShiftRows)

	if round.Round != 10 {
		fmt.Printf("After MixColumns:")
		PrintState(round.AfterMixColumns)
	}
	fmt.Printf("After AddRoundKey:")
	PrintState(round.AfterAddRoundKey)
}
