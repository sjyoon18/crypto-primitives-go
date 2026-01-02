package desviz

import (
	"crypto-primitives-go/symmetric/des"
	"fmt"
)

// Function TraceDiffusion prints the amount of different bits per round for two 64-bit inputs with the same 64-bit key.
func TraceDiffusion(p1, p2 uint64, key uint64) {

	subkeys := des.DebugSubkeys(key)

	ip1 := des.DebugInitialPermutation(p1)
	ip2 := des.DebugInitialPermutation(p2)

	L1 := uint32(ip1 >> 32)
	R1 := uint32(ip1 & 0xFFFFFFFF)

	L2 := uint32(ip2 >> 32)
	R2 := uint32(ip2 & 0xFFFFFFFF)

	fmt.Println("Diffusion")
	fmt.Println("-----------------------")
	fmt.Printf("Round 0: %d bits differ\n", BitChanged(ip1, ip2))

	for i := 0; i < 16; i++ {
		nextL1 := R1
		nextR1 := L1 ^ des.DebugF(R1, subkeys[i])

		nextL2 := R2
		nextR2 := L2 ^ des.DebugF(R2, subkeys[i])

		L1, R1 = nextL1, nextR1
		L2, R2 = nextL2, nextR2

		intermediate1 := (uint64(L1) << 32) | uint64(R1)
		intermediate2 := (uint64(L2) << 32) | uint64(R2)

		fmt.Printf(
			"Round %2d: %d bits differ\n",
			i+1,
			BitChanged(intermediate1, intermediate2),
		)
	}
}

// Function TraceSingleBitAvalanche utilizes function TraceDiffusion to trace the avalanche effect
// from a single bit flip at position bitPosition.
func TraceSingleBitAvalanche(p uint64, bitPosition int, key uint64) {
	TraceDiffusion(p, (p ^ (1 << bitPosition)), key)
}
