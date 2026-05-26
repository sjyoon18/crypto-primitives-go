package oracle

import (
	modes "crypto-primitives-go/symmetric/modes"
	"fmt"
)

func PrintOracleDemo(
	plaintext []byte,
	key [16]byte,
	iv [16]byte,
) {
	ciphertext := modes.EncryptCBC(plaintext, key, iv)

	OracleFunc := func(candidate []byte) bool {
		return PaddingOracle(candidate, key, iv)
	}

	recovered, err := RecoverPlaintext(ciphertext, iv, OracleFunc)
	if err != nil {
		panic(err)
	}

	fmt.Println("CBC Padding Oracle Attack Demo")
	fmt.Println("==============================")

	fmt.Printf("\nOriginal plaintext:\n%q\n", string(plaintext))

	fmt.Println("\nCiphertext blocks:")
	printBlocksHex(
		SplitBlocks(ciphertext, modes.AESBlockSize),
	)

	fmt.Printf("\nRecovered plaintext:\n%q\n", string(recovered))

	fmt.Println("\nPadding oracle plaintext recovery successful.")
}
