// 345. Reverse Vowels of a String
// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and
// upper cases, more than once.
package reversevowels

import "slices"

func ReverseVowels(s string) string {
	vowels := []rune{'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U'}
	reversed := make([]rune, len(s))

	i, j := 0, len(s)-1
	for i <= j {
		if !slices.Contains(vowels, rune(s[i])) {
			reversed[i] = rune(s[i])
			i++
			continue
		}

		if !slices.Contains(vowels, rune(s[j])) {
			reversed[j] = rune(s[j])
			j--
			continue
		}

		reversed[i], reversed[j] = rune(s[j]), rune(s[i])
		i++
		j--
	}

	return string(reversed)
}
