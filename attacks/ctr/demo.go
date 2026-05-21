package ctr

import (
	"encoding/hex"
	"fmt"
)

// Function PrintNonceReuseDemo prints a CTR nonce reuse demonstration,
// given two different plaintexts with same key and nonce.
func PrintNonceReuseDemo(
	p1 []byte,
	p2 []byte,
	key [16]byte,
	nonce [8]byte,
) {
	c1, c2, xored := DemonstrateNonceReuse(p1, p2, key, nonce)

	fmt.Println("CTR Nonce Reuse Attack Demo")
	fmt.Println("===========================")

	fmt.Printf("\nPlaintext 1: %q\n", string(p1))
	fmt.Printf("Plaintext 2: %q\n", string(p2))

	fmt.Printf("\nCiphertext 1: %s\n", hex.EncodeToString(c1))
	fmt.Printf("Ciphertext 2: %s\n", hex.EncodeToString(c2))

	fmt.Printf("\nP1 XOR P2: %s\n", hex.EncodeToString(xorByteSlices(p1, p2)))
	fmt.Printf("C1 XOR C2: %s\n", hex.EncodeToString(xored))

	fmt.Println("\nCTR nonce reuse detected.")
}
