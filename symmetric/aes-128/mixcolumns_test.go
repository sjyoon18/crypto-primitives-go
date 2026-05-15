package aes

import "testing"

// Function TestMixColumns verifies AES MixColumns
// using the official FIPS-197 example state.
func TestMixColumns(t *testing.T) {
	state := State{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	}
	expected := State{
		{0x04, 0xe0, 0x48, 0x28},
		{0x66, 0xcb, 0xf8, 0x06},
		{0x81, 0x19, 0xd3, 0x26},
		{0xe5, 0x9a, 0x7a, 0x4c},
	}

	mixColumns(&state)

	if state != expected {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			expected,
			state,
		)
	}
}

// Function TestInvMixColumns verifies that invMixColumns reverses mixColumns.
func TestInvMixColumns(t *testing.T) {
	state := State{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5},
	}

	original := state.Copy()
	mixColumns(&state)
	invMixColumns(&state)

	if state != original {
		t.Fatalf(
			"\nexpected:%x\ngot:	%x",
			original,
			state,
		)
	}
}
