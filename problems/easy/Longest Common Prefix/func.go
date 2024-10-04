package Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	j := 0

	for _, char := range strs[0] {

		for _, word := range strs {

			if len(word) < j+1 || rune(word[j]) != char {

				return strs[0][:j]

			}
		}

		j++
	}

	return strs[0]

}
