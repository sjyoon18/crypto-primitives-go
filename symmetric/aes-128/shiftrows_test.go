package aes

import "testing"

// Function TestShiftrows verifies AES ShiftRows
// using the official FIPS-197 example state.
func TestShiftRows(t *testing.T) {
	state := State{
		{0x87, 0xf2, 0x4d, 0x97},
		{0xec, 0x6e, 0x4c, 0x90},
		{0x4a, 0xc3, 0x46, 0xe7},
		{0x8c, 0xd8, 0x95, 0xa6},
	}
	expected := State{
		{0x87, 0xf2, 0x4d, 0x97},
		{0x6e, 0x4c, 0x90, 0xec},
		{0x46, 0xe7, 0x4a, 0xc3},
		{0xa6, 0x8c, 0xd8, 0x95},
	}
	shiftRows(&state)

	if state != expected {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			expected,
			state,
		)
	}
}

// Function TestInvShiftRows verifies that invShiftRows reverses shiftRows.
func TestInvShiftRows(t *testing.T) {
	state := State{
		{0x87, 0xf2, 0x4d, 0x97},
		{0xec, 0x6e, 0x4c, 0x90},
		{0x4a, 0xc3, 0x46, 0xe7},
		{0x8c, 0xd8, 0x95, 0xa6},
	}

	original := state.Copy()

	shiftRows(&state)
	invShiftRows(&state)

	if state != original {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			original,
			state,
		)
	}
}
