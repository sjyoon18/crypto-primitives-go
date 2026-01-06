package des

// Function DencryptBlock applies DES dencryption on a 64-bit ciphertext
// with the given 64-bit key and returns the 64-bit plaintext.
func DecryptBlock(block uint64, key uint64) uint64 {
	subkeys := roundSubkeys(key)
	permuted := initialPermutation(block)

	Ld := uint32(permuted >> 32)
	Rd := uint32(permuted & 0xFFFFFFFF)

	for i := 15; i >= 0; i-- {
		nextLd := Rd
		nextRd := Ld ^ fFunction(Rd, subkeys[i])

		Ld = nextLd
		Rd = nextRd
	}
	return finalPermutation(uint64(Rd)<<32 | uint64(Ld))
}
