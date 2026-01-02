package desviz

import "fmt"

// Function Bits64 returns the string of a 64-bit input.
func Bits64(x uint64) string {
	return fmt.Sprintf("%064b", x)
}

// Function Bits32 returns the string of a 32-bit input.
func Bits32(x uint32) string {
	return fmt.Sprintf("%032b", x)
}
