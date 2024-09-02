package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("should add an expense with a description and amount", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := Expenses{}

		expense := Expense{Amount: amount, Description: description}

		// When
		expenses.AddExpense(expense)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal(20, expenses[0].Amount)
		asserts.Equal("Lunch", expenses[0].Description)
		asserts.NotNil(expenses[0].CreatedAt)
		asserts.Nil(expenses[0].UpdatedAt)
	})

	t.Run("should add two expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := Expenses{}

		expense1 := Expense{Amount: amount, Description: description}
		expense2 := Expense{Amount: amount, Description: description}

		// When
		expenses.AddExpense(expense1)
		expenses.AddExpense(expense2)

		// Then
		asserts.Equal(2, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal(2, expenses[1].Id)
		asserts.True(expenses[0].CreatedAt.Before(expenses[1].CreatedAt))
	})

}
