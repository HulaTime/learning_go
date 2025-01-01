package popcount

// pc[i] is the poplation count of i
var pc [256]byte

// byte is an alias for uint8 in go, so this is storing 256 bytes / ints

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		/*
		  this works by working through each position in the array, and setting the population count 1
		  by 1. Starting at i=0, we set pc[0] (i/2 is 0/0). This is done by doing a bitwise AND check on
		  the righmost bit in the byte for the current index, 0 = 0b00000000, 1 = 0b00000001,
		  0&1 = 0b00000000 or 0. So the population count becomes 0.
		  Next index is 1, so we set pc[1] = pc[1/2] (1/2 is 0 here for indexing, so we get the previous
		  index), + the byte representation of 1&1, 0b00000001 AND 0b00000001 -> 0b00000001 or 1.
		  So pc[1] = pc[0] + 1 => popcount of 1
		  So pc[2] = pc[1] + 2&1 (0) => popcount of 1
		  So pc[3] = pc[2] + 3&1 (1) => popcount of 2
		  ...
		*/
	}
}

/*
  Extracting Bytes Using Shifting:
  byte(x >> (k * 8))
  This expression extracts a specific byte from the 64-bit integer x. Since x is uint64, it has
  8 bytes (8 * 8 bits = 64 bits).
    - x >> (k * 8) shifts the bits of x right by k * 8 bits.
    - byte(...) converts the result to an 8-bit value (the least significant byte after shifting).
    - The shift extracts one byte at a time.
*/

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
