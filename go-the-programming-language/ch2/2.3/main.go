package popcount

// pc[i] is the poplation count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var bytecount byte

	for i := 0; i < 8; i++ {
		bytecount += pc[byte(x>>(i*8))]
	}
	return int(bytecount)
}
