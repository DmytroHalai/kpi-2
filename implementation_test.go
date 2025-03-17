package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	t.Run("Simple expressions", func(t *testing.T) {
		res, err := PostfixToInfix("3 4 +")
		assert.NoError(t, err)
		assert.Equal(t, "(3 + 4)", res)

		res, err = PostfixToInfix("5 1 -")
		assert.NoError(t, err)
		assert.Equal(t, "(5 - 1)", res)

		res, err = PostfixToInfix("0,123 0,456 +")
		assert.NoError(t, err)
		assert.Equal(t, "(0,123 + 0,456)", res)

		res, err = PostfixToInfix("0.123 0.456 +")
		assert.NoError(t, err)
		assert.Equal(t, "(0.123 + 0.456)", res)
	})

	t.Run("Complex expressions", func(t *testing.T) {
		res, err := PostfixToInfix("2 3 1 * + 9 -")
		assert.NoError(t, err)
		assert.Equal(t, "((2 + (3 * 1)) - 9)", res)

		res, err = PostfixToInfix("3 4 5 * 6 + -")
		assert.NoError(t, err)
		assert.Equal(t, "(3 - ((4 * 5) + 6))", res)
	})

	t.Run("Very complex expressions", func(t *testing.T) {
		res, err := PostfixToInfix("123 11 3 4 + 5 6 * - 7 8 + * 9 / - ^")
		assert.NoError(t, err)
		assert.Equal(t, "(123 ^ (11 - ((((3 + 4) - (5 * 6)) * (7 + 8)) / 9)))", res)
	  })
	
	t.Run("Invalid input", func(t *testing.T) {
		_, err := PostfixToInfix("3")
		assert.Error(t, err)

		_, err = PostfixToInfix("3 + 4")
		assert.Error(t, err)

		_, err = PostfixToInfix("3 4 5 +") 
		assert.Error(t, err)
	})

	t.Run("Invalid numbers", func(t *testing.T) {
		_, err := PostfixToInfix("3 4 0,00,1123 +")
		assert.Error(t, err)

		_, err = PostfixToInfix("3 abc +")
		assert.Error(t, err)
	})

	t.Run("Number formatting", func(t *testing.T) {
		res, err := PostfixToInfix("0003 004 +")
		assert.NoError(t, err)
		assert.Equal(t, "(3 + 4)", res)
	})
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("3 4 + 2 * 7 /")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
	// Output: (((3 + 4) * 2) / 7)
}
