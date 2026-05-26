package oracle

import (
	"bytes"
	modes "crypto-primitives-go/symmetric/modes"
	"testing"
)

// Function TestRecoverPlaintext verifies that the padding oracle attack
// recovers CBC plaintext with arbitrary length.
func TestRecoverPlaintext(t *testing.T) {
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

	plaintext := []byte("CBC plaintext with arbitrary length")

	ciphertext := modes.EncryptCBC(plaintext, testKey, iv)

	oracleFunc := func(candidate []byte) bool {
		return PaddingOracle(candidate, testKey, iv)
	}

	recovered, err := RecoverPlaintext(ciphertext, iv, oracleFunc)

	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !bytes.Equal(recovered[:], plaintext) {
		t.Fatalf(
			"\nexpected: %q\ngot:	%q",
			plaintext,
			recovered,
		)
	}
}
