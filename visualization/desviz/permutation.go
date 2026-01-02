package desviz

import "fmt"

func DisplayPermutation(input uint64, table []int) {
	inBits := Bits64(input)
	for i, pos := range table {
		fmt.Printf(
			"Output bit %2d ← Input bit %2d (%c)\n",
			i+1,
			pos,
			inBits[pos-1],
		)
	}
}
