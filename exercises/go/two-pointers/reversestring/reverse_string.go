// Reverse String
//
// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.
package reversestring

func ReverseStringInPlace(s []byte) {
	i, j := 0, len(s)-1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
