package aes

import "math/bits"

// Type StateDifference represents the difference between
// two AES states.
type StateDifference struct {
	BitDifferences  int
	ByteDifferences int

	ChangedBytes [4][4]bool
	ChangedBits  [4][4]uint8
}

// Type RoundDifference represents the transformation-level
// diffusion differences occurred in a single AES round.
type RoundDifference struct {
	Round int

	StartState       StateDifference
	AfterSubBytes    StateDifference
	AfterShiftRows   StateDifference
	AfterMixColumns  StateDifference
	AfterAddRoundKey StateDifference
}

// Type EncryptionTraceDifference stores the differences
// between two AES encryption traces.
type EncryptionTraceDifference struct {
	PlaintextDifference StateDifference

	InitialAddRoundKeyDifference StateDifference
	Rounds                       []RoundDifference
	CiphertextDifference         StateDifference
}

// Function DiffStates takes two AES states and returns
// their StateDifference.
func DiffStates(a, b State) StateDifference {
	var diff StateDifference

	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			x := a[r][c] ^ b[r][c]

			if x != 0 {
				diff.ByteDifferences++
				diff.ChangedBytes[r][c] = true
			}

			bitDiffs := bits.OnesCount8(x)

			diff.BitDifferences += bitDiffs
			diff.ChangedBits[r][c] = uint8(bitDiffs)
		}
	}

	return diff
}

// Function CompareTraces traces the difference
// between two AES encryption round-by-round.
func CompareTraces(a EncryptionTrace, b EncryptionTrace) EncryptionTraceDifference {
	difference := EncryptionTraceDifference{
		PlaintextDifference:          DiffStates(a.Plaintext, b.Plaintext),
		InitialAddRoundKeyDifference: DiffStates(a.InitialAddRoundKey, b.InitialAddRoundKey),
	}

	for i := 0; i < len(a.Rounds); i++ {
		ar := a.Rounds[i]
		br := b.Rounds[i]

		rdiff := RoundDifference{
			Round: ar.Round,

			StartState:       DiffStates(ar.StartState, br.StartState),
			AfterSubBytes:    DiffStates(ar.AfterSubBytes, br.AfterSubBytes),
			AfterShiftRows:   DiffStates(ar.AfterShiftRows, br.AfterShiftRows),
			AfterAddRoundKey: DiffStates(ar.AfterAddRoundKey, br.AfterAddRoundKey),
		}

		if ar.Round != 10 {
			rdiff.AfterMixColumns = DiffStates(ar.AfterMixColumns, br.AfterMixColumns)
		}

		difference.Rounds = append(difference.Rounds, rdiff)
	}
	difference.CiphertextDifference = DiffStates(a.Ciphertext, b.Ciphertext)

	return difference
}
