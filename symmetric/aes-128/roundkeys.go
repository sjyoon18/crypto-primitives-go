package aes

// Type Word represents the 32-bit (4-byte) word in the AES.
type Word [4]byte

// Variable rCoefficient is the list of round coefficients (RC) in the AES.
var rCoefficient = [10]byte{
	0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1B, 0x36,
}

// Function rotWord returns the result of cyclically left-shifting the input word.
func rotWord(w Word) Word {
	return Word{w[1], w[2], w[3], w[0]}
}

// Function subWord returns the result of substituting each bytes in the input word
// according to the AES S-box.
func subWord(w Word) Word {
	return Word{
		sBox[w[0]],
		sBox[w[1]],
		sBox[w[2]],
		sBox[w[3]],
	}
}

// Function g represents the function g in an AES round.
// It consists rotWord, subWord, and XOR with the round coefficients.
func g(w Word, round int) Word {
	w = rotWord(w)
	w = subWord(w)
	w[0] ^= rCoefficient[round-1]
	return w
}

// Function xorWord computes and returns the XOR of two words.
func xorWord(a, b Word) Word {
	return Word{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
	}
}

// Function roundKeys returns the 44 round key words by expanding the input 16-byte key.
// This is used across both encryption and decryption rounds.
func roundKeys(key [16]byte) [44]Word {
	var w [44]Word

	for i := 0; i < 4; i++ {
		copy(w[i][:], key[4*i:4*(i+1)])
	}
	for i := 4; i < 44; i++ {
		temp := w[i-1]
		if i%4 == 0 {
			temp = g(temp, i/4)
		}
		w[i] = xorWord(w[i-4], temp)
	}
	return w
}
