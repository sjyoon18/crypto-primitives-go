package aes

import (
	"math/rand"
)

type AvalancheStatistics struct {
	NumSamples int

	AverageCiphertextDiff float64

	MinCiphertextDiff int
	MaxCiphertextDiff int

	AverageRoundDiffusion [11]float64
}

type avalancheAccumulator struct {
	totalCiphertextDiff int

	minCiphertextDiff int
	maxCiphertextDiff int

	roundTotals [11]int
}

// Function AvalancheExperiment introduces a random bit flip to the plaintext
// repeatedly to measure the AES diffusion behavior.
func AvalancheExperiment(
	plaintext [16]byte,
	key [16]byte,
	numSamples int,
) AvalancheStatistics {
	trace := EncryptWithTrace(plaintext, key)
	acc := avalancheAccumulator{
		minCiphertextDiff: 128,
	}

	for i := 0; i < numSamples; i++ {
		flipped := plaintext

		// Bit flip at random bit position [0, 127]
		bitPos := rand.Intn(128)

		flipped[bitPos/8] ^= 1 << (bitPos % 8)

		flippedTrace := EncryptWithTrace(flipped, key)

		diff := CompareTraces(trace, flippedTrace)

		finalBitDiff := diff.CiphertextDifference.BitDifferences
		acc.totalCiphertextDiff += finalBitDiff

		if finalBitDiff < acc.minCiphertextDiff {
			acc.minCiphertextDiff = finalBitDiff
		}
		if finalBitDiff > acc.maxCiphertextDiff {
			acc.maxCiphertextDiff = finalBitDiff
		}

		// Round 0
		acc.roundTotals[0] += diff.InitialAddRoundKeyDifference.BitDifferences

		// Round 1-10
		for r, round := range diff.Rounds {
			acc.roundTotals[r+1] += round.AfterAddRoundKey.BitDifferences
		}
	}

	result := AvalancheStatistics{
		NumSamples: numSamples,

		AverageCiphertextDiff: float64(acc.totalCiphertextDiff) / float64(numSamples),
		MinCiphertextDiff:     acc.minCiphertextDiff,
		MaxCiphertextDiff:     acc.maxCiphertextDiff,
	}

	for i := 0; i < 11; i++ {
		result.AverageRoundDiffusion[i] = float64(acc.roundTotals[i]) / float64(numSamples)
	}

	return result
}
