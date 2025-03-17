package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// PostfixToInfix converts a mathematical expression from postfix
// (Reverse Polish) notation to infix notation.
//
// Parameters:
//   - input: a string containing a mathematical expression in postfix
//     notation, where operands and operators are separated by spaces.
//
// Returns:
//   - A string representing the equivalent infix expression, where operators
//     are placed between operands and enclosed in parentheses to preserve operation order.
//   - An error if the input is invalid (e.g., not enough operands for an operator).
func PostfixToInfix(input string) (string, error) {
	if !hasOperands(input) {
		return "", fmt.Errorf("invalid input: %s", input)
	}
	var stack []string
	parts := strings.Fields(input)
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
			if !checkNum(part) {
				return "", fmt.Errorf("invalid number: %s", part)
			}
			stack = append(stack, cleanNumber(part))
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid input")
	}
	return stack[0], nil
}

func checkNum(input string) bool {
	if _, err := strconv.ParseFloat(strings.Replace(input, ",", ".", 1), 64); err != nil {
		return false
	}
	return true
}

func cleanNumber(input string) string {
	if strings.ContainsAny(input, ".,") {
		return "0" + strings.TrimLeft(input, "0")
	}
	return strings.TrimLeft(input, "0")
}

func hasOperands(input string) bool {
	return strings.ContainsAny(input, "+-*/^")
}
