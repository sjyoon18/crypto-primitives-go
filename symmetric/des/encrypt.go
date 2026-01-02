package des

// Function EncryptBlock returns the encrypted 64-bit input block with the 64-bit key.
func EncryptBlock(block uint64, key uint64) uint64 {
	subkeys := roundSubkeys(key)
	permuted := initialPermutation(block)

	L := uint32(permuted >> 32)
	R := uint32(permuted & 0xFFFFFFFF)

	for i := 0; i < 16; i++ {
		nextL := R
		nextR := L ^ fFunction(R, subkeys[i])

		L = nextL
		R = nextR
	}
	return finalPermutation(uint64(R)<<32 | uint64(L))
}
