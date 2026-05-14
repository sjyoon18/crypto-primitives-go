package aes

// Function computeSingleColumn takes a single 4-byte column
// and performs the AES MixColumn matrix multiplication.
// Matrix multiplication and addition is done in GF(2⁸).
func computeSingleColumn(column *[4]byte) {
	a := *column

	// Matrix multiplication
	column[0] = galoisMul(a[0], 2) ^ galoisMul(a[1], 3) ^ a[2] ^ a[3]
	column[1] = a[0] ^ galoisMul(a[1], 2) ^ galoisMul(a[2], 3) ^ a[3]
	column[2] = a[0] ^ a[1] ^ galoisMul(a[2], 2) ^ galoisMul(a[3], 3)
	column[3] = galoisMul(a[0], 3) ^ a[1] ^ a[2] ^ galoisMul(a[3], 2)
}

// Function mixColumns linearly transforms each column of the input state matrix
// in accordance with the AES MixColumns.
func mixColumns(state *State) {
	for c := 0; c < 4; c++ {
		referenceCol := [4]byte{
			state[0][c],
			state[1][c],
			state[2][c],
			state[3][c],
		}

		computeSingleColumn(&referenceCol)

		state[0][c] = referenceCol[0]
		state[1][c] = referenceCol[1]
		state[2][c] = referenceCol[2]
		state[3][c] = referenceCol[3]
	}
}

// Function invComputeSingleColumn takes a single 4-byte column
// and performs the inverse of the AES MixColumn matrix multiplication.
func invComputeSingleColumn(col *[4]byte) {
	a := *col

	// Matrix multiplication
	col[0] = galoisMul(a[0], 14) ^ galoisMul(a[1], 11) ^ galoisMul(a[2], 13) ^ galoisMul(a[3], 9)
	col[1] = galoisMul(a[0], 9) ^ galoisMul(a[1], 14) ^ galoisMul(a[2], 11) ^ galoisMul(a[3], 13)
	col[2] = galoisMul(a[0], 13) ^ galoisMul(a[1], 9) ^ galoisMul(a[2], 14) ^ galoisMul(a[3], 11)
	col[3] = galoisMul(a[0], 11) ^ galoisMul(a[1], 13) ^ galoisMul(a[2], 9) ^ galoisMul(a[3], 14)
}

// Function invMixColumns applies the inverse AES MixColumns
// to each column of the input state matrix.
func invMixColumns(state *State) {
	for c := 0; c < 4; c++ {
		referenceCol := [4]byte{
			state[0][c],
			state[1][c],
			state[2][c],
			state[3][c],
		}

		invComputeSingleColumn(&referenceCol)

		state[0][c] = referenceCol[0]
		state[1][c] = referenceCol[1]
		state[2][c] = referenceCol[2]
		state[3][c] = referenceCol[3]
	}
}

// Function MixSingleColumns returns the MixColumns result
// of the single 4-byte AES column input.
func MixSingleColumn(column [4]byte) [4]byte {
	mixed := column
	computeSingleColumn(&mixed)
	return mixed
}
