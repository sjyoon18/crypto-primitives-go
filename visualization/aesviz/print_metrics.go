package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

// Function PrintTraceDiffusionMetrics displays the bit diffusion progression
// throughout the AES encryption trace comparison.
func PrintTraceDiffusionMetrics(metrics aes.TraceDiffusionMetrics) {
	printTitle("AES Diffusion Metrics")

	fmt.Printf(
		"Plaintext difference:			%3d bits (%.2f%%)\n",
		metrics.PlaintextBits,
		metrics.PlaintextRatio*100,
	)

	fmt.Printf(
		"After initial AddRoundKey:		%3d bits (%.2f%%)\n\n",
		metrics.InitialAddRoundKeyBits,
		metrics.InitialAddRoundKeyRatio*100,
	)

	for _, round := range metrics.Rounds {
		printRoundHeader(round.Round)
		fmt.Printf(

			"  Start:          %3d bits (%.2f%%)\n",
			round.StartBits,
			round.StartRatio*100,
		)
		fmt.Printf(
			"  AfterSubBytes:  %3d bits (%.2f%%)\n",
			round.AfterSubBytesBits,
			round.AfterSubBytesRatio*100,
		)
		fmt.Printf(
			"  AfterShiftRows: %3d bits (%.2f%%)\n",
			round.AfterShiftRowsBits,
			round.AfterShiftRowsRatio*100,
		)
		if round.Round != 10 {
			fmt.Printf(
				"  AfterMixColumns:%3d bits (%.2f%%)\n",
				round.AfterMixColumnsBits,
				round.AfterMixColumnsRatio*100,
			)
		}
		fmt.Printf(
			"  AfterAddRoundKey:%3d bits (%.2f%%)\n",
			round.AfterAddRoundKeyBits,
			round.AfterAddRoundKeyRatio*100,
		)
		fmt.Println()
	}
	fmt.Printf(
		"Ciphertext difference:			%3d bits (%.2f%%)\n",
		metrics.CiphertextBits,
		metrics.CiphertextRatio*100,
	)
}
