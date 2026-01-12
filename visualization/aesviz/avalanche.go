package aesviz

import (
	aes "crypto-primitives-go/symmetric/aes-128"
	"fmt"
	"math/bits"
)

// Function HammingDistance returns the number of different bits between AES state a and b.
// This is done by counting the ones in a XOR b.
func HammingDistance(a, b aes.State) int {
	count := 0
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			count += bits.OnesCount8(a[r][c] ^ b[r][c])
		}
	}
	return count
}

// Function Print Avalanche is a helper function that prints the round and bits difference.
func PrintAvalanche(round int, bits int) {
	fmt.Printf("Round %2d: %3d bits differ\n", round, bits)
}

// Function SingleBitFlipAvalanche traces the avalanche effect
// of a single bit flip at position bitPosition.
func SingleBitFlipAvalanche(p [16]byte, bitPosition int, key [16]byte) {
	p2 := p
	p2[bitPosition/8] ^= 1 << (bitPosition % 8)

	s1 := aes.BytesToState(p)
	s2 := aes.BytesToState(p2)

	wArr := aes.DebugKeySchedule(key)
	w := wArr[:]

	aes.AddRoundKey(&s1, w, 0)
	aes.AddRoundKey(&s2, w, 0)

	PrintAvalanche(0, HammingDistance(s1, s2))

	for round := 1; round <= 9; round++ {
		aes.DebugSubBytes(&s1)
		aes.DebugSubBytes(&s2)

		aes.DebugShiftRows(&s1)
		aes.DebugShiftRows(&s2)

		aes.DebugMixColumns(&s1)
		aes.DebugMixColumns(&s2)

		aes.AddRoundKey(&s1, w, round)
		aes.AddRoundKey(&s2, w, round)

		PrintAvalanche(round, HammingDistance(s1, s2))
	}

	aes.DebugSubBytes(&s1)
	aes.DebugSubBytes(&s2)

	aes.DebugShiftRows(&s1)
	aes.DebugShiftRows(&s2)

	aes.AddRoundKey(&s1, w, 10)
	aes.AddRoundKey(&s2, w, 10)

	PrintAvalanche(10, HammingDistance(s1, s2))
}

// Function DisplaySingleBitInfluence displays the avalanche effect
// of single bit flips at all positions.
func DisplaySingleBitInfluence(p [16]byte, key [16]byte) {
	base := aes.BytesToState(aes.EncryptBlock(p, key))

	for bit := 0; bit < 128; bit++ {
		p2 := p
		p2[bit/8] ^= 1 << (bit % 8)

		alt := aes.BytesToState(aes.EncryptBlock(p2, key))
		diff := HammingDistance(base, alt)

		fmt.Printf("Bit flip at %3d: %3d output bits changed\n", bit, diff)
	}
}
