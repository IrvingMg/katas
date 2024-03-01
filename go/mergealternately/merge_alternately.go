// 1768. Merge Strings Alternately
// You are given two strings word1 and word2. Merge the strings by adding letters in
// alternating order, starting with word1. If a string is longer than the other, append the
// additional letters onto the end of the merged string.
//
// Return the merged string.
package mergealternately

func MergeAlternately(word1 string, word2 string) string {
	var merged string

	var index int
	for index < len(word1) && index < len(word2) {
		merged += string(word1[index]) + string(word2[index])
		index++
	}

	if len(word1) > index {
		merged += word1[index:]
	} else if len(word2) > index {
		merged += word2[index:]
	}

	return merged
}
