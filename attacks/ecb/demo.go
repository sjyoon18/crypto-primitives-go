package ecb

import (
	modes "crypto-primitives-go/symmetric/modes"
	"encoding/hex"
	"fmt"
)

func printBlocksASCII(blocks [][]byte) {
	for i, block := range blocks {
		fmt.Printf("[%02d] %q\n", i, string(block))
	}
}

func printBlocksHex(blocks [][]byte) {
	for i, block := range blocks {
		fmt.Printf("[%02d] %s\n", i, hex.EncodeToString(block))
	}
}

// Function PrintECBLeakageDemo prints plaintext blocks, ciphertext blocks,
// and repeated-block leakage for the ECB mode.
func PrintECBLeakageDemo(
	plaintext []byte,
	key [16]byte,
) {
	ciphertext, repeated := DemonstrateECBLeakage(plaintext, key)

	fmt.Println("ECB Repeated Block Leakage")
	fmt.Println("==========================")

	fmt.Println("Plaintext blocks:")
	printBlocksASCII(
		SplitBlocks(plaintext, modes.AESBlockSize),
	)

	fmt.Println("\nCiphertext blocks:")
	printBlocksHex(
		SplitBlocks(ciphertext, modes.AESBlockSize),
	)

	fmt.Println("\nLeakage result:")

	if len(repeated) == 0 {
		fmt.Println("No repeated ciphertext blocks detected.")
		return
	}

	for _, block := range repeated {
		fmt.Printf(
			"Repeated block detected %d times: %s\n",
			block.Count,
			hex.EncodeToString(block.Block),
		)
	}
}
