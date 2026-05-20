package ecb

import (
	modes "crypto-primitives-go/symmetric/modes"
	"testing"
)

// Function TestDetectRepeatedBlocks verifies that repeated ciphertext blocks are detected.
func TestDetectRepeatedBlocks(t *testing.T) {
	data := []byte(
		"SixteenByteBlock" + "SixteenByteBlock",
	)

	repeated := DetectRepeatedBlocks(data, modes.AESBlockSize)

	if len(repeated) != 1 {
		t.Fatalf("expected 1 repeated block, got %d repeated block(s)", len(repeated))
	}

	if repeated[0].Count != 2 {
		t.Fatalf(
			"expected repeated block count 2, got %d repeated block count(s)",
			repeated[0].Count,
		)
	}
}

// Function TestDetectNoRepeatedBlocks verifies that unique blocks
// are not identified as repeated.
func TestDetectNoRepeatedBlocks(t *testing.T) {
	data := []byte(
		"SixteenByteBlock" + "DifferentBlock!!",
	)

	repeated := DetectRepeatedBlocks(data, modes.AESBlockSize)

	if len(repeated) != 0 {
		t.Fatalf("expected no repeated blocks, got %d repeated blocks", len(repeated))
	}
}

// Function TestDemonstrateECBLeakage verifies that ECB mode leaks repeated blocks.
func TestDemonstrateECBLeakage(t *testing.T) {
	data := []byte(
		"SixteenByteBlock" + "SixteenByteBlock" + "SixteenByteBlock",
	)

	testKey := [16]byte{
		0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0x0b,
		0x0c, 0x0d, 0x0e, 0x0f,
	}

	_, repeated := DemonstrateECBLeakage(data, testKey)

	if len(repeated) == 0 {
		t.Fatal("ECB repeated-block leakage expected")
	}
}
