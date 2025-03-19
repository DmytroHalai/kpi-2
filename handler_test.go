package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHandler(t *testing.T) {
	t.Run("EmptyInput", func(t *testing.T) {
		input := strings.NewReader("")
		output := &strings.Builder{}

		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "empty row")
	})

	t.Run("InvalidExpression", func(t *testing.T) {
		input := strings.NewReader("3 +")
		output := &strings.Builder{}

		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input")
	})
}

func ExampleComputeHandler_Compute() {
	var input io.Reader
	var output io.Writer = os.Stdout

	handler := ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
