package pascal_s_triangle

// generateBruteForce
//
// Given an integer numRows, return the first numRows of Pascal's triangle.
//
// In Pascal's triangle, each number is the sum of the two numbers directly above
func generateBruteForce(numRows int) (triangle [][]int) {
	triangle = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		lenI := i + 1
		triangle[i] = make([]int, lenI)
		if i == 0 {
			triangle[i][0] = 1
			continue
		}

		for ii := 0; ii < lenI; ii++ {
			leftI := ii - 1
			if leftI < 0 {
				triangle[i][ii] = triangle[i-1][ii]
				continue
			}
			rightI := ii
			if rightI >= i {
				triangle[i][ii] = triangle[i-1][ii-1]
				continue
			}
			triangle[i][ii] = triangle[i-1][leftI] + triangle[i-1][rightI]
		}
	}
	return
}

// generateAlternative
//
// NOT FINISHED!
//
// The idea is simple:
// it seems that all the values either predictable or calculable
//
// i=0            1                  1
// i=1           1|1                 1|1
// i=2          1|2|1                1|i|1
// i=3         1|3|3|1               1|i|i|1
// i=4        1|4|6|4|1              1|i|i*((i+1)/3)+i/2|i  |1
// i=5      1|5|10|10|5|1            1|i|   i*i/2       |...
// i=6     1|6|15|20|15|6|1          1|i|i*((i+1)/3)+i/2| i*3+(i/3)|...
// i=7    1|7|21|35|35|21|7|1        1|i|   i*i/2       |...
// i=8   1|8|28|56|70|56|28|8|1      1|i|i*((i+1)/3)+i/2| i*(i-1) |...
// i=9  1|9|36|84|126|126|84|36|9|   1|i|   i*i/2       |...
//
// And it seems like there's a way to find a universal formula for all the values
func generateAlternative(numRows int) (triangle [][]int) {
	triangle = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		lenI := i + 1
		triangle[i] = make([]int, lenI)
		for ii := 0; ii < lenI/2+lenI%2; ii++ {
			if i == 0 {
				triangle[i][ii] = 1
				continue
			}
			if ii == 0 {
				triangle[i][ii], triangle[i][i-ii] = 1, 1
				//triangle[i][0], triangle[i][1] = 1, 1
				continue
			}
			if ii == 1 {
				triangle[i][ii], triangle[i][i-ii] = i, i
				continue
			}
			if ii == 2 {
				var k int
				if i%2 == 1 {
					k = i * (i / 2)
				} else {
					k = i*((i+1)/3) + i/2
				}
				triangle[i][ii], triangle[i][i-ii] = k, k
				continue
			}
			//TODO: finalize it
		}
	}
	return
}
