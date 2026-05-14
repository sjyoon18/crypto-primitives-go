package aes

// Function addRoundKeyState XORs the AES state with the round key.
func addRoundKeyState(state *State, roundKey State) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			state[r][c] ^= roundKey[r][c]
		}
	}
}
