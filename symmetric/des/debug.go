package des

// Function DebugInitialPermutation exposes the IP for visualization.
func DebugInitialPermutation(x uint64) uint64 {
	return initialPermutation(x)
}

// Function DebugF exposes the f function for visualization.
func DebugF(r uint32, k uint64) uint32 {
	return fFunction(r, k)
}

// Function DebugSubkeys exposes the subkeys for visualization.
func DebugSubkeys(key uint64) [16]uint64 {
	return roundSubkeys(key)
}
