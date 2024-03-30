// Package number_of_one_bits
//
// Write a function that takes the binary representation of a positive integer and returns the number of
// set bits
// it has (also known as the Hamming weight).
package number_of_one_bits

// iterativeTill0
// Iteratively we increment the counter, if the last bit of the number is 0,
// then we shift the input number right and check if the number is 0.
func iterativeTill0(n int) (number int) {
	for n != 0 {
		number += n & 1
		n = n >> 1
	}
	return
}
