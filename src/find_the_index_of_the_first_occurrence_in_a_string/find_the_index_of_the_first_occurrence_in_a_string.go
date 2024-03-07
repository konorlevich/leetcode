package find_the_index_of_the_first_occurrence_in_a_string

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	for i := range haystack {
		if haystackLen < needleLen {
			return -1
		}
		if haystack[i:i+needleLen] == needle {
			return i
		}
		haystackLen--
	}
	return -1
}
