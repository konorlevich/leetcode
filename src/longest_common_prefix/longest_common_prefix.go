package longest_common_prefix

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
