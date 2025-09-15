// Valid Parentheses
//
// Description:
// Write a function that takes a string of parentheses, and determines if the
// order of the parentheses is valid. The function should return true if the
// string is valid, and false if it's invalid.
//
// Examples:
//
//	"()"              =>  true
//	")(()))"          =>  false
//	"("               =>  false
//	"(())((()())())"  =>  true
package parenvalidator

func Validate(str string) bool {
	var balance int

	for _, s := range str {
		if s == '(' {
			balance++
			continue
		}

		balance--
		if balance < 0 {
			break
		}
	}

	return balance == 0
}

func ValidateRecursively(str string) bool {
	var balance int

	validateRecursively(str, &balance)

	return balance == 0
}

func validateRecursively(str string, balance *int) {
	if *balance < 0 || len(str) == 0 {
		return
	}

	if str[0] == '(' {
		*balance++
	} else {
		*balance--
	}
	validateRecursively(str[1:], balance)
}
