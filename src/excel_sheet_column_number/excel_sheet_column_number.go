// Package excel_sheet_column_number
//
// Given a string columnTitle that represents the column title as appears in an Excel sheet,
// return its corresponding column number.
//
// For example:
//
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
//
// Constraints:
//
// 1 <= columnTitle.length <= 7
//
// columnTitle consists only of uppercase English letters.
//
// columnTitle is in the range ["A", "FXSHRXW"].
package excel_sheet_column_number

// titleToNumber
//
// We use simple ASCII-math.
// ASCII A is 65, ASCII Z is 90.
// We iterate over the string from the last symbol to the first (why not?).
// For each symbol we calculate its value by subtracting 64.
// We then multiply this value by the variability (26), to a power equal to the symbol position (from the end)
func titleToNumber(columnTitle string) int {
	lettersCount := 26
	var number = 0
	if len(columnTitle) == 0 {
		return number
	}
	lastIndex := len(columnTitle) - 1
	for i := lastIndex; i >= 0; i-- {
		symbol := columnTitle[i]
		num := int(symbol) - 64 // ASCII A = 65, in our task A = 1
		if i == lastIndex {     // it is unnecessary, but it makes the func a bit faster
			number += num
			continue
		}
		digitNumber := lastIndex - i
		number += num * pow(lettersCount, digitNumber)
	}

	return number
}

func pow(x, y int) (powed int) {
	if x == 1 || y == 0 {
		return 1
	}
	powed = x
	if y == 1 {
		return powed
	}
	for i := 1; i < y; i++ {
		powed *= x
	}
	return
}
