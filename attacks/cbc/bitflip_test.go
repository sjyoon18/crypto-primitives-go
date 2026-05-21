package cbc

import (
	"bytes"
	modes "crypto-primitives-go/symmetric/modes"
	"testing"
)

// Function TestCBCBitFlip verifies that modifying the previous ciphertext block
// causes a desired change in the next decrypted plaintext block.
func TestCBCBitFlip(t *testing.T) {
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

	plaintext := []byte(
		"userID:xxxxxxxx;" +
			"balance:0000000;",
	)

	ciphertext := modes.EncryptCBC(plaintext, testKey, iv)

	modified := FlipCiphertext(ciphertext, 1, 8, []byte("0000000;"), []byte("9999999;"))

	decrypted, err := modes.DecryptCBC(modified, testKey, iv)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !bytes.Contains(decrypted, []byte("balance:9999999;")) {
		t.Fatalf(
			"expected decrypted plaintext: balance:9999999;\ngot: %q",
			string(decrypted),
		)
	}
}
