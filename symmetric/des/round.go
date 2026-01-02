package des

// Function fFunction returns the result of applying the DES Feistel function to the 32-bit input.
func fFunction(r uint32, subkey uint64) uint32 {
	expand := expansionE(r)
	xor := expand ^ subkey
	sboxed := sBoxReplace(xor)
	return fPermutation(sboxed)
}
