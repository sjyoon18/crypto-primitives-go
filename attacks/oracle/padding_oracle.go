package oracle

import modes "crypto-primitives-go/symmetric/modes"

// Function PaddingOracle evaluates whether the ciphertext has valid PKCS#7 padding
// and returns a boolean accordingly.
func PaddingOracle(
	ciphertext []byte,
	key [16]byte,
	iv [16]byte,
) bool {
	_, err := modes.DecryptCBC(ciphertext, key, iv)
	return err == nil
}
