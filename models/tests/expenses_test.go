package tests

import (
	"expense-tracker/models"
	"expense-tracker/models/tests/dsl"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("✅ should add an expense with a description and amount", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}

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

	t.Run("✅ should add two expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense1 := models.Expense{Amount: amount, Description: description}
		expense2 := models.Expense{Amount: amount, Description: description}

		// When
		expenses.AddExpense(expense1)
		expenses.AddExpense(expense2)

		// Then
		asserts.Equal(2, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal(2, expenses[1].Id)
		asserts.True(expenses[0].CreatedAt.Before(expenses[1].CreatedAt))
	})

	t.Run("✅ should update an expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		// When
		expense.Id = 1
		expense.Description = "Dinner"
		expenses.UpdateExpense(expense)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal("Dinner", expenses[0].Description)
		asserts.True(expenses[0].CreatedAt.Before(*expenses[0].UpdatedAt))
	})

	t.Run("❌ should not updated an non-existent expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		// When
		expense.Id = 2
		expense.Description = "Dinner"
		expenses.UpdateExpense(expense)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal("Lunch", expenses[0].Description)
		asserts.Nil(expenses[0].UpdatedAt)
	})

	t.Run("✅ should delete an expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		// When
		expenses.DeleteExpense(1)

		// Then
		asserts.Equal(0, len(expenses))
	})

	t.Run("❌ should not delete an non-existent expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		// When
		expenses.DeleteExpense(2)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal("Lunch", expenses[0].Description)
		asserts.Nil(expenses[0].UpdatedAt)
	})

	t.Run("✅ should list all expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		expectedMessage := dsl.JoinMessage(expenses)
		expectedMessage = fmt.Sprintf(models.HeaderFormat + "\n" + expectedMessage)

		// When
		result := dsl.OutputToString(expenses.ListExpenses)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal("Lunch", expenses[0].Description)
		asserts.Nil(expenses[0].UpdatedAt)
		asserts.Equal(expectedMessage, result)
	})

	t.Run("✅ should print the header when there are no expenses", func(t *testing.T) {
		// Given
		expenses := models.Expenses{}

		// When
		result := dsl.OutputToString(expenses.ListExpenses)

		// Then
		asserts.Equal(models.HeaderFormat+"\n", result)
	})

	t.Run("✅ should print the summary of all expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		amount = 15
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.AddExpense(expense)

		// When
		result := dsl.OutputToString(expenses.Summary)

		// Then
		asserts.Equal("Total expenses: 35\n", result)
	})

	// should print the summary of all expenses for a specific month of the current year
}
