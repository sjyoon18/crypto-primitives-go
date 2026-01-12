package aes

// Function DebugKeySchedule exposes the round keys for visualization.
func DebugKeySchedule(key [16]byte) [44]Word {
	return roundKeys(key)
}

// Function DebugSubBytes exposes the substituted state for visualization.
func DebugSubBytes(state *State) {
	subBytes(state)
}

// Function DebugShiftRows exposes the shiftRows operation for visualization.
func DebugShiftRows(state *State) {
	shiftRows(state)
}

// Function DebugMixColumns exposes the mixColumns operation for visualization.
func DebugMixColumns(state *State) {
	mixColumns(state)
}
