// 1071. Greatest Common Divisor of Strings
// For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t +
// t (i.e., t is concatenated with itself one or more times).
//
// Given two strings str1 and str2, return the largest string x such that x divides both
// str1 and str2.
package gcdofstrings

import "strings"

func GcdOfStrings(str1 string, str2 string) string {
	if str1 == str2 {
		return str1
	}

	if len(str2) > len(str1) {
		return GcdOfStrings(str2, str1)
	}

	after, found := strings.CutPrefix(str1, str2)
	if found {
		return GcdOfStrings(after, str2)
	}

	return ""
}
