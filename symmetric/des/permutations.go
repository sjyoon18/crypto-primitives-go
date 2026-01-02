package des

// ipTable is the permutation table for the initial permutation in DES.
var ipTable = []int{
	58, 50, 42, 34, 26, 18, 10, 2,
	60, 52, 44, 36, 28, 20, 12, 4,
	62, 54, 46, 38, 30, 22, 14, 6,
	64, 56, 48, 40, 32, 24, 16, 8,
	57, 49, 41, 33, 25, 17, 9, 1,
	59, 51, 43, 35, 27, 19, 11, 3,
	61, 53, 45, 37, 29, 21, 13, 5,
	63, 55, 47, 39, 31, 23, 15, 7,
}

// ipInverseTable is the permutation table for the final permutation in DES.
var ipInverseTable = []int{
	40, 8, 48, 16, 56, 24, 64, 32,
	39, 7, 47, 15, 55, 23, 63, 31,
	38, 6, 46, 14, 54, 22, 62, 30,
	37, 5, 45, 13, 53, 21, 61, 29,
	36, 4, 44, 12, 52, 20, 60, 28,
	35, 3, 43, 11, 51, 19, 59, 27,
	34, 2, 42, 10, 50, 18, 58, 26,
	33, 1, 41, 9, 49, 17, 57, 25,
}

// permute is a generic function that returns a 64-bit output by permuting a 64-bit input.
// The permutation is done according to the given permTable.
// inputSize specifies how many bits are meaningful from the 64-bit input.
func permute(input uint64, permTable []int, inputSize int) uint64 {
	var output uint64 = 0

	for i, pos := range permTable {
		bit := (input >> (inputSize - pos)) & 1
		output |= bit << (len(permTable) - 1 - i)
	}
	return output
}

// Function initialPermutation returns a 64-bit output by permuting the 64-bit input block.
// Initial permutation IP is done according to the ipTable.
func initialPermutation(block uint64) uint64 {
	return permute(block, ipTable, 64)
}

// Function finalPermutation returns a 64-bit output by permuting the 64-bit input block.
// Final permutation IP-1 is done according to the ipInverseTable and is the inverse operation of IP.
func finalPermutation(block uint64) uint64 {
	return permute(block, ipInverseTable, 64)
}
