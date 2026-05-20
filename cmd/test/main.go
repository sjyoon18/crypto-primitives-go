package main

import ecbattack "crypto-primitives-go/attacks/ecb"

func main() {
	key := [16]byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x15, 0x88,
		0x09, 0xcf, 0x4f, 0x3c,
	}

	plaintext := []byte(
		"SixteenByteBlock" +
			"SixteenByteBlock" +
			"DifferentBlock!!" +
			"SixteenByteBlock",
	)

	ecbattack.PrintECBLeakageDemo(plaintext, key)
}
