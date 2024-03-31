// Package longest_common_prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
package longest_common_prefix

// longestCommonPrefix
func longestCommonPrefix(strs []string) (result string) {
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i <= 200; i++ {
		s := ""
		cur := ""
		for _, str := range strs {
			if len(str) <= i {
				return result
			}
			if s == "" {
				s = str[i : i+1]
				continue
			}
			cur = str[i : i+1]
			if cur != s {
				return result
			}
		}
		result += cur
	}
	return result
}
