package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
)

func SingleBitFlipTraceDifference(

	plaintext [16]byte,
	bitPosition int,
	key [16]byte,

) aes.EncryptionTraceDifference {

	modified := plaintext
	modified[bitPosition/8] ^= 1 << (bitPosition % 8)
	traceA := aes.EncryptWithTrace(plaintext, key)
	traceB := aes.EncryptWithTrace(modified, key)
	return aes.CompareTraces(traceA, traceB)

}

func PrintSingleBitFlipAvalanche(

	plaintext [16]byte,
	bitPosition int,
	key [16]byte,

) {

	diff := SingleBitFlipTraceDifference(
		plaintext,
		bitPosition,
		key,
	)
	fmt.Printf("AES single-bit avalanche trace\n")
	fmt.Printf("Flipped plaintext bit: %d\n", bitPosition)
	fmt.Println("================================")
	fmt.Printf(
		"After initial AddRoundKey: %d bits differ\n",
		diff.InitialAddRoundKeyDifference.BitDifferences,
	)
	for _, round := range diff.Rounds {
		fmt.Printf(
			"Round %2d final state: %3d bits differ\n",
			round.Round,
			round.AfterAddRoundKey.BitDifferences,
		)
	}
	fmt.Printf(
		"Ciphertext: %d bits differ\n",
		diff.CiphertextDifference.BitDifferences,
	)

}

func DisplaySingleBitInfluence(

	plaintext [16]byte,
	key [16]byte,

) {

	baseTrace := aes.EncryptWithTrace(plaintext, key)
	for bit := 0; bit < 128; bit++ {
		modified := plaintext
		modified[bit/8] ^= 1 << (bit % 8)
		modifiedTrace := aes.EncryptWithTrace(modified, key)
		diff := aes.CompareTraces(
			baseTrace,
			modifiedTrace,
		)
		fmt.Printf(
			"Bit flip at %3d: %3d output bits changed\n",
			bit,
			diff.CiphertextDifference.BitDifferences,
		)
	}

}
