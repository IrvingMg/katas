// Reverse Array
//
// Description:
// Write a function that reverses each string in an array. The input is given
// as an array of strings.
//
// Examples:
//
//	Input: s = ["alpha", "beta", "charlie", "delta", "foxtrot"]
//	Output: ["ahpla", "ateb", "eilrahc", "atled", "tortxof"]
package reversearray

import (
	"slices"
)

func ReverseMultiString(data []string) []string {
	for i := range data {
		data[i] = reverseString(data[i])
	}

	return data
}

func reverseString(str string) string {
	runes := []rune(str)
	slices.Reverse(runes)

	return string(runes)
}
