package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

func PrintAvalancheStatistics(stats aes.AvalancheStatistics) {
	fmt.Println("AES Avalanche Statistics")
	fmt.Println("==========================")

	fmt.Printf(
		"Number of samples: %d\n",
		stats.NumSamples,
	)

	fmt.Printf(
		"Average ciphertext diffusion: %.2f bits\n",
		stats.AverageCiphertextDiff,
	)

	fmt.Printf(
		"Minimum ciphertext diffusion: %d bits\n",
		stats.MinCiphertextDiff,
	)

	fmt.Printf(
		"Maximum ciphertext diffusion: %d bits\n",
		stats.MaxCiphertextDiff,
	)

	fmt.Println("\nAverage round diffusion:")

	for i, val := range stats.AverageRoundDiffusion {
		if i == 0 {
			fmt.Printf(
				"Round 0 (Initial AddRoundKey): %.2f bits\n",
				val,
			)
		} else {
			fmt.Printf(
				"Round %-2d: %.2f bits\n",
				i,
				val,
			)
		}
	}
}
