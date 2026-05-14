package aes

// Function EncryptBlock applies AES-128 encryption on a 16-byte plaintext
// with the given 16-byte key and returns the 16-byte ciphertext.
func EncryptBlock(plaintext [16]byte, key [16]byte) [16]byte {
	state := bytesToState(plaintext)
	rKeys := roundKeys(key)

	// Round 0
	addRoundKeyState(&state, wordsToState(rKeys[0:4]))

	// Round 1 - 9
	for round := 1; round <= 9; round++ {
		subBytes(&state)
		shiftRows(&state)
		mixColumns(&state)
		addRoundKeyState(&state, wordsToState(rKeys[round*4:(round+1)*4]))
	}

	// Final round
	subBytes(&state)
	shiftRows(&state)
	addRoundKeyState(&state, wordsToState(rKeys[40:44]))

	return stateToBytes(state)
}
