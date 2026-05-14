package aes

// Function ShiftRows cyclically right-shifts each state row.
// Row 0: unchanged.
// Row 1: three bytes to the right.
// Row 2: two bytes to the right.
// Row 3: one byte to the right.
func shiftRows(state *State) {
	for r := 1; r < 4; r++ {
		state[r] = rightRotate(state[r], 4-r)
	}
}

// Function invShiftRows cyclically shifts each state row inversely.
func invShiftRows(state *State) {
	for r := 1; r < 4; r++ {
		state[r] = rightRotate(state[r], r)
	}
}

// Function rightRotate returns the result of cyclically right-shifting a 4-byte row by n positions.
func rightRotate(row [4]byte, n int) [4]byte {
	var result [4]byte
	for i := 0; i < 4; i++ {
		result[(i+n)%4] = row[(i)]
	}
	return result
}
