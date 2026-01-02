package main

import (
	"crypto-primitives-go/visualization/desviz"
	"fmt"
)

func main() {
	key := uint64(0x133457799BBCDFF1)
	plaintext := uint64(0x0123456789ABCDEF)
	flipPosition := 63
	manipulated := plaintext ^ (uint64(1) << uint64(flipPosition))

	fmt.Println("DES Diffusion Visualization")
	fmt.Println("=============================")
	fmt.Printf("Plaintext 1: %064b\n", plaintext)
	fmt.Printf("Plaintext 2: %064b\n\n", manipulated)

	desviz.TraceDiffusion(plaintext, manipulated, key)

	fmt.Println("\nOutcome comparison (avalanche effect):")
	fmt.Println("=========================================")
	desviz.DisplaySingleBitFlipAvalanche(plaintext, key)
}
