package ctr

import modes "crypto-primitives-go/symmetric/modes"

// Function xorByteSlices returns the byte-wise XOR of two byte slices up to the shorter length.
func xorByteSlices(b1, b2 []byte) []byte {
	minLen := min(len(b1), len(b2))

	out := make([]byte, minLen)

	for i := 0; i < minLen; i++ {
		out[i] = b1[i] ^ b2[i]
	}
	return out
}

// Function DemonstrateNonceReuse encrypts two plaintexts P1, P2 with the same nonce and key.
// It returns the ciphertexts C1, C2, and the recovered value of P1 XOR P2 from C1 XOR C2.
func DemonstrateNonceReuse(
	p1 []byte,
	p2 []byte,
	key [16]byte,
	nonce [8]byte,
) ([]byte, []byte, []byte) {
	c1 := modes.EncryptCTR(p1, key, nonce)
	c2 := modes.EncryptCTR(p2, key, nonce)

	recoveredXOR := xorByteSlices(c1, c2)

	return c1, c2, recoveredXOR
}
