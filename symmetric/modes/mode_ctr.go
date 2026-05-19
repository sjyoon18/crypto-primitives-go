package modes

import aes "crypto-primitives-go/symmetric/aes-128"

// Function buildCounterBlock builds a 16-byte AES input counter block
// from an 8-byte nonce and an 8-byte counter.
func buildCounterBlock(nonce [8]byte, counter uint64) [16]byte {
	var block [16]byte

	copy(block[0:8], nonce[:])

	for i := 0; i < 8; i++ {
		block[15-i] = byte(counter >> (8 * i))
	}

	return block
}

// Function incrementCounter increments the counter portion (last 64 bits)
// of the 16-byte counterBlock.
func incrementCounter(counterBlock *[16]byte) {
	for i := 15; i >= 8; i-- {
		counterBlock[i]++
		if counterBlock[i] != 0 {
			break
		}
	}
}

// Function ApplyCTR is the root function for encryption and decryption
// using AES-128 in CTR mode.
func ApplyCTR(
	input []byte,
	key [16]byte,
	nonce [8]byte,
) []byte {
	output := make([]byte, len(input))

	counterBlock := buildCounterBlock(nonce, 0)

	for i := 0; i < len(input); i += AESBlockSize {
		keystream := aes.EncryptBlock(counterBlock, key)

		blockSize := AESBlockSize
		if i+blockSize > len(input) {
			blockSize = len(input) - i
		}

		for j := 0; j < blockSize; j++ {
			output[i+j] = input[i+j] ^ keystream[j]
		}

		incrementCounter(&counterBlock)
	}

	return output
}

// Function EncryptCTR encrypts plaintext in AES-128 CTR mode.
func EncryptCTR(
	plaintext []byte,
	key [16]byte,
	nonce [8]byte,
) []byte {
	return ApplyCTR(plaintext, key, nonce)
}

// Function DecryptCTR decrypts ciphertext in AES-128 CTR mode.
func DecryptCTR(
	ciphertext []byte,
	key [16]byte,
	nonce [8]byte,
) []byte {
	return ApplyCTR(ciphertext, key, nonce)
}
