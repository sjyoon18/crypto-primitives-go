package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

type HeatmapOptions struct {
	Symbols    bool
	ShowTotals bool
}

func heatSymbol(n uint8) string {
	switch {
	case n == 0:
		return "☐"
	case n <= 2:
		return "░"
	case n <= 5:
		return "▒"
	default:
		return "█"
	}
}

// Function PrintBitHeatmap takes a StateDifference and prints
// per-byte bit diffusion in a 4 x 4 AES state difference matrix.
func PrintBitHeatmap(
	diff aes.StateDifference,
	option HeatmapOptions,
) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if option.Symbols {
				fmt.Printf("%s", heatSymbol(diff.ChangedBits[r][c]))
			} else {
				fmt.Printf("%2d ", diff.ChangedBits[r][c])
			}
		}
		fmt.Println()
	}
}

// Function PrintRoundHeatmap prints AES diffusion heatmaps for each round.
func PrintRoundHeatmaps(
	traceDiff aes.EncryptionTraceDifference,
	options HeatmapOptions,
) {
	printTitle("AES Round Diffusion Heatmaps")

	printRoundHeader(0)
	printSection("After Initial AddRoundKey:")
	PrintBitHeatmap(traceDiff.InitialAddRoundKeyDifference, options)

	for _, roundDiff := range traceDiff.Rounds {
		printRoundHeader(roundDiff.Round)

		printSection("After SubBytes:")
		PrintBitHeatmap(roundDiff.AfterSubBytes, options)

		printSection("After ShiftRows:")
		PrintBitHeatmap(roundDiff.AfterShiftRows, options)

		if roundDiff.Round != 10 {
			printSection("After MixColumns")
			PrintBitHeatmap(roundDiff.AfterMixColumns, options)
		}

		printSection("After AddRoundKey:")
		PrintBitHeatmap(roundDiff.AfterAddRoundKey, options)
	}
}
