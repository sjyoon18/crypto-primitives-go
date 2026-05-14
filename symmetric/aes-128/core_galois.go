package aes

// Function galoisMul returns the result of multiplying two bytes a and b in GF(2⁸).
// The irreducible polynomial used for AES is x⁸ + x⁴ + x³ + x + 1 (0x11b).
func galoisMul(a, b byte) byte {
	var c byte
	for i := 0; i < 8; i++ {
		if (b & 1) != 0 {
			c ^= a
		}

		ms := a & 0x80
		a <<= 1

		if ms != 0 {
			a ^= 0x1b
		}

		b >>= 1

	}
	return c
}
