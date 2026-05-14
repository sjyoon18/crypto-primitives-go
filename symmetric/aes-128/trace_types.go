package aes

type RoundTrace struct {
	Round int

	StartState State

	AfterSubBytes    State
	AfterShiftRows   State
	AfterMixColumns  State
	AfterAddRoundKey State

	RoundKey State
}

type EncryptionTrace struct {
	Plaintext State

	InitialAddRoundKey State

	Rounds []RoundTrace

	Ciphertext State
}

func (s State) Copy() State {
	var out State

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			out[r][c] = s[r][c]
		}
	}
	return out
}
