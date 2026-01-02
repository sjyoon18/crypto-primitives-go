package des

// eTable is the table of the expansion permutation E in DES.
var eTable = []int{
	32, 1, 2, 3, 4, 5, 4, 5, 6, 7, 8, 9, 8, 9, 10, 11, 12, 13, 12, 13, 14, 15, 16, 17,
	16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1,
}

// pTable is the table of the permutation P within the f function in DES.
var pTable = []int{
	16, 7, 20, 21,
	29, 12, 28, 17,
	1, 15, 23, 26,
	5, 18, 31, 10,
	2, 8, 24, 14,
	32, 27, 3, 9,
	19, 13, 30, 6,
	22, 11, 4, 25,
}

// Function expansionE returns a 48-bit output by expanding the 32-bit input.
// This is done by bit swaps according to the eTable.
func expansionE(right uint32) uint64 {
	return permute(uint64(right), eTable, 32)
}

// Function fPermutation returns a 32-bit output by permuting the 32-bit input.
// Permutation is done according to the pTable.
func fPermutation(x uint32) uint32 {
	return uint32(permute(uint64(x), pTable, 32))
}
