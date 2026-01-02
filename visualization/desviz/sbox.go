package desviz

import "fmt"

func VisualizeSBox(input uint8, sbox [4][16]uint8) {
	row := ((input & 0b100000) >> 4) | (input & 0b000001)
	col := (input >> 1) & 0b1111
	output := sbox[row][col]

	fmt.Printf("Input:  %06b\n", input)
	fmt.Printf("Row:    %02b (%d)\n", row, row)
	fmt.Printf("Column: %04b (%d)\n", col, col)
	fmt.Printf("Output: %04b\n", output)
}
