package cbc

import (
	modes "crypto-primitives-go/symmetric/modes"
	"fmt"
)

// Function PrintCBCBitFlipDemo prints a CBC bit-flip attack demonstration.
func PrintCBCBitFlipDemo(
	plaintext []byte,
	key [16]byte,
	iv [16]byte,
	blockIndex int,
	offset int,
	original []byte,
	desired []byte,
) {
	ciphertext := modes.EncryptCBC(plaintext, key, iv)

	modified := FlipCiphertext(
		ciphertext,
		blockIndex,
		offset,
		original,
		desired,
	)

	modifiedPlaintext, err := modes.DecryptCBC(modified, key, iv)

	if err != nil {
		panic(err)
	}

	fmt.Println("CBC Bit-Flipping Attack Demo")
	fmt.Println("============================")

	fmt.Printf("\nOriginal plaintext:\n%q\n", string(plaintext))

	fmt.Printf(
		"\nDesired modification:\n%q ==> %q\n",
		string(original),
		string(desired),
	)

	fmt.Println("\nUnmodified ciphertext blocks:")
	printBlocksHex(
		SplitBlocks(ciphertext, modes.AESBlockSize),
	)

	fmt.Println("\nModified ciphertext blocks:")
	printBlocksHex(
		SplitBlocks(modified, modes.AESBlockSize),
	)

	fmt.Println("\nModified plaintext blocks (%q):")
	printBlocksQuoted(SplitBlocks(modifiedPlaintext, modes.AESBlockSize))

	fmt.Println("\nModified plaintext blocks (%s):")
	printBlocksString(SplitBlocks(modifiedPlaintext, modes.AESBlockSize))

	fmt.Println("\nCBC plaintext modification successful.")
}
