package modes

import aes "crypto-primitives-go/symmetric/aes-128"

// Function EncryptCBC encrypts a plaintext with arbitrary length
// using AES-128 in CBC mode with PKCS#7 padding.
func EncryptCBC(plaintext []byte, key [16]byte, iv [16]byte) []byte {
	padded := PadPKCS7(plaintext, AESBlockSize)

	ciphertext := make([]byte, len(padded))

	previous := iv[:]

	for i := 0; i < len(padded); i += AESBlockSize {
		block := padded[i : i+AESBlockSize]

		xored := xorBlocks(block, previous)

		var aesBlock [16]byte
		copy(aesBlock[:], xored)

		encrypted := aes.EncryptBlock(aesBlock, key)

		copy(ciphertext[i:i+AESBlockSize], encrypted[:])

		previous = ciphertext[i : i+AESBlockSize]
	}

	return ciphertext
}

// Function DecryptCBC decrypts AES-128 CBC ciphertext and removes
// the PKCS#7 padding.
func DecryptCBC(ciphertext []byte, key [16]byte, iv [16]byte) ([]byte, error) {
	if len(ciphertext)%AESBlockSize != 0 {
		return nil, ErrInvalidCiphertextLength
	}

	plaintext := make([]byte, len(ciphertext))

	previous := iv[:]

	for i := 0; i < len(ciphertext); i += AESBlockSize {
		var block [16]byte
		copy(block[:], ciphertext[i:i+AESBlockSize])

		decrypted := aes.DecryptBlock(block, key)

		xored := xorBlocks(decrypted[:], previous)

		copy(plaintext[i:i+AESBlockSize], xored)

		previous = ciphertext[i : i+AESBlockSize]
	}

	return UnpadPKCS7(plaintext, AESBlockSize)
}
