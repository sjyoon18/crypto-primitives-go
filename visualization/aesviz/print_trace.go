package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
)

func PrintTrace(trace aes.EncryptionTrace) {
	printTitle("AES-128 Encryption Trace")

	printSection("Plaintext State:")
	PrintState(trace.Plaintext)

	printSection("After Initial Key Addition:")
	PrintState(trace.InitialAddRoundKey)

	for _, round := range trace.Rounds {
		PrintRoundTrace(round)
	}

	printSection("Ciphertext State:")
	PrintState(trace.Ciphertext)
}

func PrintRoundTrace(round aes.RoundTrace) {
	printRoundHeader(round.Round)

	printSection("Start State:")
	PrintState(round.StartState)

	printSection("Round Key:")
	PrintState(round.RoundKey)

	printSection("After SubBytes:")
	PrintState(round.AfterSubBytes)

	printSection("After ShiftRows:")
	PrintState(round.AfterShiftRows)

	if round.Round != 10 {
		printSection("After MixColumns:")
		PrintState(round.AfterMixColumns)
	}
	printSection("After AddRoundKey:")
	PrintState(round.AfterAddRoundKey)
}
