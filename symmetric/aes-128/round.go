package aes

// Function addRoundKeyState XORs the AES state with the round key.
func addRoundKeyState(state *State, roundKey State) {
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			state[r][c] ^= roundKey[r][c]
		}
	}
}

// Function AddRoundKey XORs the AES state with the round key in a given round.
func AddRoundKey(state *State, word []Word, round int) {
	roundKey := wordsToState(word[4*round : 4*(round+1)])
	addRoundKeyState(state, roundKey)
}
