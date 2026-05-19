package modes

import (
	"bytes"
	"testing"
)

// Function TestEncryptDecryptCTR verifies that CTR decryption reverses CTR encryption.
func TestEncryptDecryptCTR(t *testing.T) {
	plaintext := []byte("test message: hello CTR mode!")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	nonce := [8]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
	}

	ciphertext := EncryptCTR(plaintext, testKey, nonce)
	got := DecryptCTR(ciphertext, testKey, nonce)

	if !bytes.Equal(got, plaintext) {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			plaintext,
			got,
		)
	}
}

// Function TestCTRShortBlock verifies that CTR mode preserves the plaintext length
// without padding.
func TestCTRShortBlock(t *testing.T) {
	plaintext := []byte("short")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	nonce := [8]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
	}

	ciphertext := EncryptCTR(plaintext, testKey, nonce)

	if len(ciphertext) != len(plaintext) {
		t.Fatalf(
			"expected ciphertext length %d, got %d",
			len(plaintext),
			len(ciphertext),
		)
	}
}

// Function TestCTRDeterministic verifies that with a given plaintext CTR mode produces
// the same ciphertext for the same key and nonce.
func TestCTRDeterministic(t *testing.T) {
	plaintext := []byte("same plaintext and same nonce")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	nonce := [8]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
	}

	c1 := EncryptCTR(plaintext, testKey, nonce)
	c2 := EncryptCTR(plaintext, testKey, nonce)

	if !bytes.Equal(c1, c2) {
		t.Fatal("different ciphertext is produced with same nonce")
	}
}

// Function TestCTRDDifferentNonce verifies that with a given plaintext CTR mode produces
// different ciphertexts for the same key and different nonce.
func TestCTRDifferentNonce(t *testing.T) {
	plaintext := []byte("same plaintext and different nonce")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	n1 := [8]byte{
		0x10, 0x11, 0x12, 0x13,
		0x14, 0x15, 0x16, 0x17,
	}

	n2 := [8]byte{
		0x20, 0x21, 0x22, 0x23,
		0x24, 0x25, 0x26, 0x27,
	}

	c1 := EncryptCTR(plaintext, testKey, n1)
	c2 := EncryptCTR(plaintext, testKey, n2)

	if bytes.Equal(c1, c2) {
		t.Fatal("same ciphertext is produced with different nonce")
	}
}
