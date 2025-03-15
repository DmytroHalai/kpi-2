package lab2

import (
	"fmt"
	"io"
	"strings"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

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
