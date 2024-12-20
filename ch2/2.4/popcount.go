package popcount

func PopCount(x uint64) int {
	var bytecount byte

	// shift for each bit in the 64 bit x input
	var popcount uint64 = 0
	for i := 0; i < 64; i++ {
		popcount += x & 1
		x >>= 1
	}
	return int(bytecount)
}
