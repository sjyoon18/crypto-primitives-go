package desviz

import (
	"crypto-primitives-go/symmetric/des"
	"fmt"
	"math/bits"
)

// Function HammingDistance returns the number of different bits between 64-bit inputs a and b.
// This is done by counting the ones in a XOR b.
func HammingDistance(a, b uint64) int {
	return bits.OnesCount64(a ^ b)
}

// Function SingleBitFlipAvalanche returns the number of bit changed
// from a single bit flip in the 64-bit input at position bitPosition.
// bitPosition starts from the MSB.
func SingleBitFlipAvalanche(p uint64, bitPosition int, key uint64) int {
	p2 := p ^ (uint64(1) << (63 - bitPosition))

	c1 := des.EncryptBlock(p, key)
	c2 := des.EncryptBlock(p2, key)

	return HammingDistance(c1, c2)
}

// Function DisplaySingleBitFlipAvalanche utilizes function SingleBitFlipAvalanche
// to trace the avalanche effect of single bit flips at all positions.
func DisplaySingleBitFlipAvalanche(p uint64, key uint64) {
	for i := 0; i < 64; i++ {
		fmt.Printf("Bit flip at position %2d: %2d bits changed\n",
			i+1,
			SingleBitFlipAvalanche(p, i, key))
	}
}
