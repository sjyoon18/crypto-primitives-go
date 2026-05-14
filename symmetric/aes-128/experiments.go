package aes

import "math/rand"

// Type SingleBitFlipExperiment stores the AES trace with the original plaintext
// and that with the bit-flipped modified plaintext.
type SingleBitFlipExperiment struct {
	BitPosition int

	OriginalPlaintext [16]byte
	ModifiedPlaintext [16]byte
	Key               [16]byte

	OriginalTrace EncryptionTrace
	ModifiedTrace EncryptionTrace

	Difference EncryptionTraceDifference
}

// Type AvalancheStatistics stores aggregate avalanche metrics
// from multiple single-bit flip experiments
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

// Function FlipPlaintextBit returns the modified version of the plaintext
// with bit flip at bitPosition [0, 127].
func FlipPlaintextBit(
	plaintext [16]byte,
	bitPosition int,
) [16]byte {
	if bitPosition < 0 || bitPosition >= 128 {
		panic("bitPosition out of range [0, 127]")
	}

	modified := plaintext
	modified[bitPosition/8] ^= 1 << (bitPosition % 8)

	return modified
}

func RunSingleBitFlipExperiment(
	plaintext [16]byte,
	key [16]byte,
	bitPosition int,
) SingleBitFlipExperiment {
	modified := FlipPlaintextBit(
		plaintext,
		bitPosition,
	)
	originalTrace := EncryptWithTrace(plaintext, key)
	modifiedTrace := EncryptWithTrace(modified, key)

	diff := CompareTraces(originalTrace, modifiedTrace)

	return SingleBitFlipExperiment{
		BitPosition: bitPosition,

		OriginalPlaintext: plaintext,
		ModifiedPlaintext: modified,
		Key:               key,

		OriginalTrace: originalTrace,
		ModifiedTrace: modifiedTrace,

		Difference: diff,
	}
}

// Function AvalancheExperiment introduces a random bit flip to the plaintext
// repeatedly to measure the AES diffusion behavior.
func AvalancheExperiment(
	plaintext [16]byte,
	key [16]byte,
	numSamples int,
) AvalancheStatistics {

	if numSamples <= 0 {
		panic("numSamples must be greater than 0")
	}

	acc := avalancheAccumulator{
		minCiphertextDiff: 128,
	}

	for i := 0; i < numSamples; i++ {
		// Bit flip at random bit position [0, 127]
		bitPosition := rand.Intn(128)
		exp := RunSingleBitFlipExperiment(
			plaintext,
			key,
			bitPosition,
		)

		diff := exp.Difference

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
