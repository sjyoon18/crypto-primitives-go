package cbc

import (
	"encoding/hex"
	"fmt"
)

func printBlocksHex(blocks [][]byte) {
	for i, block := range blocks {
		fmt.Printf("[%02d] %s\n", i, hex.EncodeToString(block))
	}
}

func printBlocksQuoted(blocks [][]byte) {
	for i, block := range blocks {
		fmt.Printf("[%02d] %q\n", i, string(block))
	}

}

func printBlocksString(blocks [][]byte) {
	for i, block := range blocks {
		fmt.Printf("[%02d] %s\n", i, string(block))
	}

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
