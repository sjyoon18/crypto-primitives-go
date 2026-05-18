package modes

// Function xorBlocks returns the XOR value of two byte blocks.
func xorBlocks(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("blocks must have equal length")
	}
	result := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}

	return result
}
