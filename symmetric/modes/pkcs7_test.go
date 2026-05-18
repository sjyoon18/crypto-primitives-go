package modes

import "testing"

func TestPadPKCS7(t *testing.T) {
	input := []byte("plaintext message")

	expected := []byte{
		'p', 'l', 'a', 'i',
		'n', 't', 'e', 'x',
		't', ' ', 'm', 'e',
		's', 's', 'a', 'g',
		'e', 0x03, 0x03, 0x03,
	}

	got := PadPKCS7(input, 20)

	if string(got) != string(expected) {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			expected,
			got,
		)
	}
}

func TestUnpadPKCS7(t *testing.T) {
	input := []byte{
		'p', 'l', 'a', 'i',
		'n', 't', 'e', 'x',
		't', ' ', 'm', 'e',
		's', 's', 'a', 'g',
		'e', 0x03, 0x03, 0x03,
	}

	expected := []byte("plaintext message")

	got, err := UnpadPKCS7(input, 20)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if string(got) != string(expected) {
		t.Fatalf(
			"\nexpected: %x\ngot:	%x",
			expected,
			got,
		)
	}
}
