package oracle

import modes "crypto-primitives-go/symmetric/modes"

// Type OracleFunc represents a padding oracle which returns
// true if the ciphertext input has valid padding and false otherwise.
type OracleFunc func(ciphertext []byte) bool

// Function twoBlockCiphertext builds the two-block ciphertext used for
// one padding oracle query.
func twoBlockCiphertext(previousBlock [16]byte, targetBlock [16]byte) []byte {
	out := make([]byte, 0, 2*modes.AESBlockSize)
	out = append(out, previousBlock[:]...)
	out = append(out, targetBlock[:]...)
	return out
}

// Function RecoverBlock recovers the plaintext block of the targetBlock
// by modifying previousBlock and querying the oracle.
func RecoverBlock(
	previousBlock [16]byte,
	targetBlock [16]byte,
	oracle OracleFunc,
) [16]byte {
	var intermediate [16]byte
	var recovered [16]byte

	modifiedPrevious := previousBlock

	for pos := modes.AESBlockSize - 1; pos >= 0; pos-- {
		paddingValue := byte(modes.AESBlockSize - pos)

		for i := pos + 1; i < modes.AESBlockSize; i++ {
			modifiedPrevious[i] = intermediate[i] ^ paddingValue
		}

		found := false

		for guess := 0; guess < 256; guess++ {
			modifiedPrevious[pos] = byte(guess)

			if !oracle(twoBlockCiphertext(modifiedPrevious, targetBlock)) {
				continue
			}

			// Detect false positives
			if pos == modes.AESBlockSize-1 {
				check := modifiedPrevious
				check[pos-1] ^= 0x01

				if !oracle(twoBlockCiphertext(check, targetBlock)) {
					continue
				}
			}
			intermediate[pos] = byte(guess) ^ paddingValue
			recovered[pos] = intermediate[pos] ^ previousBlock[pos]
			found = true
			break
		}

		if !found {
			panic("failed to recover")
		}

	}
	return recovered
}

// Function RecoverPlaintext recovers the full CBC plaintext from ciphertext
// using the IV and a padding oracle.
func RecoverPlaintext(
	ciphertext []byte,
	iv [16]byte,
	oracle OracleFunc,
) ([]byte, error) {
	if len(ciphertext)%modes.AESBlockSize != 0 {
		return nil, modes.ErrInvalidCiphertextLength
	}

	blocks := make([][16]byte, 0)
	blocks = append(blocks, iv)

	for i := 0; i < len(ciphertext); i += modes.AESBlockSize {
		var block [16]byte
		copy(block[:], ciphertext[i:i+modes.AESBlockSize])
		blocks = append(blocks, block)
	}

	recovered := make([]byte, 0, len(ciphertext))

	for i := 1; i < len(blocks); i++ {
		blockPlaintext := RecoverBlock(blocks[i-1], blocks[i], oracle)

		recovered = append(recovered, blockPlaintext[:]...)
	}

	return modes.UnpadPKCS7(recovered, modes.AESBlockSize)
}
