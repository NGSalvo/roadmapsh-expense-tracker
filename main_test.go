package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("should an expense with a description and amount", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := Expenses{}

		expense := Expense{Amount: amount, Description: description}

		// When
		expenses.AddExpense(expense)

		// Then
		asserts.Equal(len(expenses), 1)
		asserts.Equal(expenses[0].Amount, 20)
		asserts.Equal(expenses[0].Description, "Lunch")
		asserts.NotNil(expenses[0].CreatedAt)
		asserts.Nil(expenses[0].UpdatedAt)
	})

}
