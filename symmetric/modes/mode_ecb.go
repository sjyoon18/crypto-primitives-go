package modes

import (
	aes "crypto-primitives-go/symmetric/aes-128"
)

const AESBlockSize = 16

// Function EncryptECB encrypts a plaintext with arbitrary length
// using AES-128 in ECB mode with PKCS#7 padding.
func EncryptECB(plaintext []byte, key [16]byte) []byte {
	padded := PadPKCS7(plaintext, AESBlockSize)

	ciphertext := make([]byte, len(padded))

	for i := 0; i < len(padded); i += AESBlockSize {
		var block [16]byte
		copy(block[:], padded[i:i+AESBlockSize])

		encrypted := aes.EncryptBlock(block, key)

		copy(ciphertext[i:i+AESBlockSize], encrypted[:])
	}

	return ciphertext
}

// Function DecryptECB decrypts AES-128 ECB ciphertext and removes
// the PKCS#7 padding.
func DecryptECB(ciphertext []byte, key [16]byte) ([]byte, error) {
	if len(ciphertext)%AESBlockSize != 0 {
		return nil, ErrInvalidCiphertextLength
	}

	plaintext := make([]byte, len(ciphertext))

	for i := 0; i < len(ciphertext); i += AESBlockSize {
		var block [16]byte
		copy(block[:], ciphertext[i:i+AESBlockSize])

		decrypted := aes.DecryptBlock(block, key)

		copy(plaintext[i:i+AESBlockSize], decrypted[:])
	}

	return UnpadPKCS7(plaintext, AESBlockSize)
}
