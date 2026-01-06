package desviz

import "fmt"

// Function DisplayPermutation visualizes the permutation according to a given table.
func DisplayPermutation(input uint64, table []int) {
	inBits := Bits64(input)
	for i, pos := range table {
		fmt.Printf(
			"Input bit %2d → Output bit %2d (bit: %c)",
			pos,
			i+1,
			inBits[pos-1],
		)
	}
}
