package ctr

import "testing"

// Function TestRecoverPlaintextXOR verifies that CTR nonce reuse leaks
// the information of P1 XOR P2.
func TestRecoverPlaintextXOR(t *testing.T) {
	p1 := []byte("plaintext one")
	p2 := []byte("plaintext two")

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

	_, _, got := DemonstrateNonceReuse(p1, p2, testKey, nonce)

	expected := xorByteSlices(p1, p2)

	if string(got) != string(expected) {
		t.Fatalf(
			"\nexpected :%x\ngot:	%x",
			expected,
			got,
		)
	}
}
