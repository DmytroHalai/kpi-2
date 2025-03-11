package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// PostfixToInfix converts
func PostfixToInfix(input string) (string, error) {
	var stack []string
	parts := strings.Split(input, " ")

	for _, part := range parts {
		switch part {
		case "+", "-", "*", "/", "^":
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid input")
			}
			op2, op1 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			newExpr := fmt.Sprintf("(%s %s %s)", op1, part, op2)
			stack = append(stack, newExpr)
		default:
			stack = append(stack, part)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid input")
	}
	return stack[0], nil
}
