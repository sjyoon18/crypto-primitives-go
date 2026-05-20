package ecb

import "crypto-primitives-go/symmetric/modes"

// Type BlockFrequency stores a ciphertext block and how many times it appeared.
type BlockFrequency struct {
	Block []byte
	Count int
}

// Function SplitBlocks splits data with arbitrary length into fixed-size blocks.
func SplitBlocks(data []byte, blockSize int) [][]byte {
	if blockSize <= 0 {
		panic("invalid block size")
	}

	blocks := make([][]byte, 0)

	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize
		if end > len(data) {
			end = len(data)
		}

		block := make([]byte, end-i)
		copy(block, data[i:end])

		blocks = append(blocks, block)
	}

	return blocks
}

// Function DetectRepeatedBlocks returns a list of BlockFrequency of the repeated blocks.
func DetectRepeatedBlocks(data []byte, blockSize int) []BlockFrequency {
	blocks := SplitBlocks(data, blockSize)

	counts := make(map[string]int)
	rawBlocks := make(map[string][]byte)

	for _, block := range blocks {
		if len(block) != blockSize {
			continue
		}

		key := string(block)
		counts[key]++
		rawBlocks[key] = block
	}

	repeated := make([]BlockFrequency, 0)

	for key, count := range counts {
		if count > 1 {
			repeated = append(repeated, BlockFrequency{
				Block: rawBlocks[key],
				Count: count,
			})
		}
	}

	return repeated
}

func HasECBLeakage(ciphertext []byte, blockSize int) bool {
	return len(DetectRepeatedBlocks(ciphertext, blockSize)) > 0
}

func DemonstrateECBLeakage(plaintext []byte, key [16]byte) ([]byte, []BlockFrequency) {
	ciphertext := modes.EncryptECB(plaintext, key)
	repeated := DetectRepeatedBlocks(ciphertext, modes.AESBlockSize)

	return ciphertext, repeated
}
