package aes

// Type State represents the internal AES state.
// It is a 128-bit data path that consists 16 bytes arranged in a 4x4 byte matrix.
type State [4][4]byte

// Function BytesToState builds and returns an AES state from the 16-byte input.
func bytesToState(b [16]byte) State {
	var s State
	for i := 0; i < 16; i++ {
		s[i%4][i/4] = b[i]
	}
	return s
}

// Function StateToBytes builds and returns a 16-byte output from the AES state input.
func stateToBytes(s State) [16]byte {
	var b [16]byte
	for i := 0; i < 16; i++ {
		b[i] = s[i%4][i/4]
	}
	return b
}

// Function wordsToState returns an AES state from four consecutive key schedule words.
func wordsToState(words []Word) State {
	var s State
	for c := 0; c < 4; c++ {
		s[0][c] = words[c][0]
		s[1][c] = words[c][1]
		s[2][c] = words[c][2]
		s[3][c] = words[c][3]
	}
	return s
}
