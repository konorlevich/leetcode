// Package reverse_bits
//
// Reverse bits of a given 32 bits unsigned integer.
//
// Constraints: The input must be a binary string of length 32
//
// Follow up: If this function is called many times, how would you optimize it?
package reverse_bits

// iterative solution
//
// We check if `nth` bit from the right side from the `num` is 1.
// We create a number 1 and shift it left to the position we are checking.
// Then we shift the resulting (reversed) number left and add the num we found shifted back right.
func iterative(num uint32) (reversed uint32) {
	for i := 0; i < 32; i++ {
		reversed = (reversed << 1) | (num&(uint32(1)<<i))>>i
	}
	return reversed
}

// iterativeWithRightShift
//
// Similar solution, but this time we shift `num` after each iteration.
//
// We check if the last bit is 1, shift reversed number left and set the last bit to the value we found.
// Then we shift `num` right.
func iterativeWithRightShift(num uint32) (reversed uint32) {
	for i := 0; i < 32; i++ {
		reversed = (reversed << 1) | (num & 1)
		num = num >> 1
	}
	return reversed
}
