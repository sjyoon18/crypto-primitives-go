package cbc

import modes "crypto-primitives-go/symmetric/modes"

// Function FlipCiphertext modifies the previous CBC ciphertext block
// to obtain a desired change in the next decrypted plaintext block.
func FlipCiphertext(
	ciphertext []byte,
	blockIndex int,
	offset int,
	original []byte,
	desired []byte,
) []byte {
	if blockIndex <= 0 {
		panic("invalid block index")
	}

	if len(original) != len(desired) {
		panic("original and desired must have equal length")
	}

	modified := make([]byte, len(ciphertext))
	copy(modified, ciphertext)

	previousBlockStart := (blockIndex - 1) * modes.AESBlockSize

	for i := 0; i < len(original); i++ {
		delta := original[i] ^ desired[i]
		modified[previousBlockStart+offset+i] ^= delta
	}

	return modified
}
