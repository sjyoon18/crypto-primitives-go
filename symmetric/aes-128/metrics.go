package aes

// Type RoundDiffusionMetric summarizes the difference of two executions
// at the end of each AES round.
type RoundDiffusionMetric struct {
	Round int

	StartBits            int
	AfterSubBytesBits    int
	AfterShiftRowsBits   int
	AfterMixColumnsBits  int
	AfterAddRoundKeyBits int

	StartRatio            float64
	AfterSubBytesRatio    float64
	AfterShiftRowsRatio   float64
	AfterMixColumnsRatio  float64
	AfterAddRoundKeyRatio float64
}

// TraceDiffusionMetrics summarizes diffusion across the full AES trace.
type TraceDiffusionMetrics struct {
	PlaintextBits          int
	InitialAddRoundKeyBits int
	CiphertextBits         int

	PlaintextRatio          float64
	InitialAddRoundKeyRatio float64
	CiphertextRatio         float64

	Rounds []RoundDiffusionMetric
}

// Function bitRatio computes the ratio of the input bit amount for 128 bits.
func bitRatio(bits int) float64 {
	return float64(bits) / 128.0
}

// Function AnalyzeTraceDiffusion converts a trace difference
// into round-by-round diffusion metrics.
func AnalyzeTraceDiffusion(diff EncryptionTraceDifference) TraceDiffusionMetrics {
	metrics := TraceDiffusionMetrics{
		PlaintextBits:          diff.PlaintextDifference.BitDifferences,
		InitialAddRoundKeyBits: diff.InitialAddRoundKeyDifference.BitDifferences,
		CiphertextBits:         diff.CiphertextDifference.BitDifferences,

		PlaintextRatio:          bitRatio(diff.PlaintextDifference.BitDifferences),
		InitialAddRoundKeyRatio: bitRatio(diff.InitialAddRoundKeyDifference.BitDifferences),
		CiphertextRatio:         bitRatio(diff.CiphertextDifference.BitDifferences),
	}

	for _, round := range diff.Rounds {
		rmetric := RoundDiffusionMetric{
			Round: round.Round,

			StartBits:            round.StartState.BitDifferences,
			AfterSubBytesBits:    round.AfterSubBytes.BitDifferences,
			AfterShiftRowsBits:   round.AfterShiftRows.BitDifferences,
			AfterMixColumnsBits:  round.AfterMixColumns.BitDifferences,
			AfterAddRoundKeyBits: round.AfterAddRoundKey.BitDifferences,

			StartRatio:            bitRatio(round.StartState.BitDifferences),
			AfterSubBytesRatio:    bitRatio(round.AfterSubBytes.BitDifferences),
			AfterShiftRowsRatio:   bitRatio(round.AfterShiftRows.BitDifferences),
			AfterMixColumnsRatio:  bitRatio(round.AfterMixColumns.BitDifferences),
			AfterAddRoundKeyRatio: bitRatio(round.AfterAddRoundKey.BitDifferences),
		}
		metrics.Rounds = append(metrics.Rounds, rmetric)
	}
	return metrics
}
