// Brackets validator
//
// Given an expression string exp, write a program to examine whether the pairs
// and the orders of “{“, “}”, “(“, “)”, “[“, “]” are correct in the given
// expression.
//
// Examples:
//
//	“[()]{}{[()()]()}”       =>  true
//	“[(])”                   =>  false
package bracketsvalidator

func Validate(exp string) bool {
	openingBrackets := make([]rune, 0)

	for _, s := range exp {
		if s == '{' || s == '[' || s == '(' {
			openingBrackets = append(openingBrackets, s)
			continue
		}

		if len(openingBrackets) == 0 {
			return false
		}

		lastElementIndex := len(openingBrackets) - 1
		lastOpenBracket := openingBrackets[lastElementIndex]
		if (s == '}' && lastOpenBracket == '{') ||
			(s == ']' && lastOpenBracket == '[') ||
			(s == ')' && lastOpenBracket == '(') {
			openingBrackets = openingBrackets[:lastElementIndex]
			continue
		}

		return false
	}

	return len(openingBrackets) == 0
}
