package aes

// Function DencryptBlock applies AES-128 dencryption on a 16-byte ciphertext
// with the given 16-byte key and returns the 16-byte plaintext.
func Decrypt(ciphertext [16]byte, key [16]byte) [16]byte {
	state := bytesToState(ciphertext)
	roundKeys := roundKeys(key)

	// Round 10
	addRoundKey(&state, wordsToState(roundKeys[40:44]))

	// Round 9 - 1
	for round := 9; round >= 1; round-- {
		invShiftRows(&state)
		invSubBytes(&state)
		addRoundKey(&state, wordsToState(roundKeys[round*4:(round+1)*4]))
		invMixColumns(&state)
	}

	// Final round
	invShiftRows(&state)
	invSubBytes(&state)
	addRoundKey(&state, wordsToState(roundKeys[0:4]))

	return stateToBytes(state)
}
