package oracle

import (
	modes "crypto-primitives-go/symmetric/modes"
	"testing"
)

// Function TestPaddingOracleValidPadding verifies that ciphertext
// with valid PKCS#7 padding is accepted.
func TestPaddingOracleValidPadding(t *testing.T) {
	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	iv := [16]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f,
	}

	plaintext := []byte("test message: padding oracle")

	ciphertext := modes.EncryptCBC(plaintext, testKey, iv)

	if !PaddingOracle(ciphertext, testKey, iv) {
		t.Fatal("expected oracle to accept valid padding")
	}
}

// Function TestPaddingOracleInvalidPadding verifies that corrupted ciphertext
// with invalid padding is rejected.
func TestPaddingOracleInvalidPadding(t *testing.T) {
	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	iv := [16]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
		0x18, 0x19, 0x1a, 0x1b,
		0x1c, 0x1d, 0x1e, 0x1f,
	}

	plaintext := []byte("test message: padding oracle")

	ciphertext := modes.EncryptCBC(plaintext, testKey, iv)

	corrupted := make([]byte, len(ciphertext))
	copy(corrupted, ciphertext)

	corrupted[len(corrupted)-1] ^= 0xff

	if PaddingOracle(corrupted, testKey, iv) {
		t.Fatal("expected oracle to reject invalid padding")
	}
}
