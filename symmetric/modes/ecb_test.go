package modes

import (
	"bytes"
	"testing"
)

// Function TestEncryptDecryptECB verifies that ECB decryption
// reverses ECB encryption.
func TestEncryptDecryptECB(t *testing.T) {
	plaintext := []byte("test message: hello ECB mode!")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	ciphertext := EncryptECB(plaintext, testKey)

	got, err := DecryptECB(ciphertext, testKey)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !bytes.Equal(got, plaintext) {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			plaintext,
			got,
		)
	}
}

// Function TestECBBlockAlignment verifies that ECB ciphertext
// is block-aligned.
func TestECBBlockAlignment(t *testing.T) {
	plaintext := []byte("test message test message")

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	ciphertext := EncryptECB(plaintext, testKey)

	if len(ciphertext)%AESBlockSize != 0 {
		t.Fatalf("ciphertext length is not block aligned")
	}

	if len(ciphertext) <= len(plaintext) {
		t.Fatalf("ciphertext must include PKCS#7 padding")
	}
}

// Function TestDecryptECBInvalidLength verifies that ECB rejects
// malformed ciphertext lengths.
func TestDecryptECBInvalidLength(t *testing.T) {
	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	invalidCiphertext := []byte{0x00, 0x01, 0x02}

	_, err := DecryptECB(invalidCiphertext, testKey)

	if err == nil {
		t.Fatal("error is expected for invalid ciphertext length")
	}
}
