// Example 1: Given a string s, return true if it is a palindrome, false otherwise.
//
// A string is a palindrome if it reads the same forward as backward.
// That means, after reversing it, it is still the same string. For example: "abcdcba", or "racecar".
package palindrome

func IsPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
