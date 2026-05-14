package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

func PrintSingleBitFlipAvalanche(exp aes.SingleBitFlipExperiment) {
	diff := exp.Difference

	fmt.Println("AES Single-Bit Avalnche Experiment")
	fmt.Println("====================================")

	fmt.Printf(
		"Flipped plaintext bit: %d\n",
		exp.BitPosition,
	)

	fmt.Printf(
		"After initial AddRoundKey: %d bits differ\n",
		diff.InitialAddRoundKeyDifference.BitDifferences,
	)

	for _, round := range diff.Rounds {
		fmt.Printf(
			"After round %2d: %3d bits differ \n",
			round.Round,
			round.AfterAddRoundKey.BitDifferences,
		)
	}

	fmt.Printf(
		"Ciphertext: %d bits differ\n",
		diff.CiphertextDifference.BitDifferences,
	)
}

func PrintSingleBitInfluence(
	plaintext [16]byte,
	key [16]byte,
) {
	for bit := 0; bit < 128; bit++ {
		exp := aes.RunSingleBitFlipExperiment(
			plaintext,
			key,
			bit,
		)

		fmt.Printf(
			"Bit flip at %3d: %3d output bits changed\n",
			bit,
			exp.Difference.CiphertextDifference.BitDifferences,
		)
	}

}
