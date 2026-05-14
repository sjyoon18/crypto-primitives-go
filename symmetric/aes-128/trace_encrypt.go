package aes

func EncryptWithTrace(plaintext [16]byte, key [16]byte) EncryptionTrace {
	state := BytesToState(plaintext)
	rKeys := roundKeys(key)

	trace := EncryptionTrace{
		Plaintext: state.Copy(),
	}

	// Round 0
	addRoundKeyState(&state, wordsToState(rKeys[0:4]))
	trace.InitialAddRoundKey = state.Copy()

	// Round 1-9
	for r := 1; r <= 9; r++ {
		rt := RoundTrace{
			Round:      r,
			StartState: state.Copy(),
			RoundKey:   wordsToState(rKeys[r*4 : (r+1)*4]),
		}

		subBytes(&state)
		rt.AfterSubBytes = state.Copy()

		shiftRows(&state)
		rt.AfterShiftRows = state.Copy()

		mixColumns(&state)
		rt.AfterMixColumns = state.Copy()

		addRoundKeyState(&state, rt.RoundKey)
		rt.AfterAddRoundKey = state.Copy()

		trace.Rounds = append(trace.Rounds, rt)
	}

	// Round 10
	final := RoundTrace{
		Round:      10,
		StartState: state.Copy(),
		RoundKey:   wordsToState(rKeys[40:44]),
	}

	subBytes(&state)
	final.AfterSubBytes = state.Copy()

	shiftRows(&state)
	final.AfterShiftRows = state.Copy()

	addRoundKeyState(&state, final.RoundKey)
	final.AfterAddRoundKey = state.Copy()

	trace.Rounds = append(trace.Rounds, final)

	trace.Ciphertext = state.Copy()

	return trace
}
