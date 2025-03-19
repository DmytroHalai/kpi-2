package lab2

import (
	"fmt"
	"io"
	"strings"
)

// ComputeHandler represents a handler that processes mathematical expressions.
//
// It contains two fields: Input and Output.
//
//   - Input is an io.Reader that is used to read the input data (mathematical expression).
//   - Output is an io.Writer where the result will be written after computation.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// Compute is a method of ComputeHandler that reads a mathematical expression from the Input,
// processes it, and writes the result to the Output.
//
// It performs the following steps:
//
//   - Reads the input expression and stores it in a buffer.
//   - Strips leading and trailing whitespace from the expression.
//   - If the expression is empty, returns an error indicating that the row is empty.
//   - Converts the postfix expression (if any) into an infix expression using PostfixToInfix function.
//   - Writes the resulting infix expression to the Output.
func (ch *ComputeHandler) Compute() error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, ch.Input)
	if err != nil {
		return fmt.Errorf("error reading: %w", err)
	}

	expression := strings.TrimSpace(buf.String())
	if expression == "" {
		return fmt.Errorf("empty row")
	}

	result, err := PostfixToInfix(expression)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(ch.Output, result)
	if err != nil {
		return fmt.Errorf("error writing: %w", err)
	}

	return nil
}
