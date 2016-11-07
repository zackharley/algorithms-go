/**
 * Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 * The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.
 */

package main

import (
	"fmt"
)

func validateParensString(str string) bool {

	if len(str) % 2 != 0 {
		return false
	}

	stackLength := len(str) / 2
	stack := make([]byte, 0, stackLength)

	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == '[' || str[i] == '{' {
			stack = append(stack, str[i])
		} else {
			if stack[len(stack) - 1] == '(' && str[i] != ')' {
				return false
			} else if stack[len(stack) - 1] == '[' && str[i] != ']' {
				return false
			} else if stack[len(stack) - 1] == '{' && str[i] != '}' {
				return false
			}

			stack = stack[0:len(stack) - 1] // Pop value off of the top of the stack
		}
 	}

	return true
}

func main() {
	fmt.Println(validateParensString("()")) // true
	fmt.Println(validateParensString("()[]{}")) // true
	fmt.Println(validateParensString("([{}])")) // true

	fmt.Println(validateParensString("{[}]")) // false
	fmt.Println(validateParensString("(]")) // false
	fmt.Println(validateParensString("[[][[[]]][][])")) // false

}