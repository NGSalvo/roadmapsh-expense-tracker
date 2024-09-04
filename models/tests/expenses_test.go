package tests

import (
	"expense-tracker/models"
	"expense-tracker/models/tests/dsl"
	"fmt"
	"testing"
	"time"

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
		err := expenses.Add(expense)

		// Then
		asserts.Equal(1, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal(20, expenses[0].Amount)
		asserts.Equal("Lunch", expenses[0].Description)
		asserts.NotNil(expenses[0].CreatedAt)
		asserts.Nil(expenses[0].UpdatedAt)
		asserts.Nil(err)
	})

	t.Run("✅ should add two expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense1 := models.Expense{Amount: amount, Description: description}
		expense2 := models.Expense{Amount: amount, Description: description}

		// When
		err := expenses.Add(expense1)
		err2 := expenses.Add(expense2)

		// Then
		asserts.Equal(2, len(expenses))
		asserts.Equal(1, expenses[0].Id)
		asserts.Equal(2, expenses[1].Id)
		asserts.True(expenses[0].CreatedAt.Before(expenses[1].CreatedAt))
		asserts.Nil(err)
		asserts.Nil(err2)
	})

	t.Run("❌ should not add an expense with a negative amount", func(t *testing.T) {
		// Given
		amount := -20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}

		// When
		expenses.Add(expense)

		// Then
		asserts.Equal(0, len(expenses))
		asserts.Error(fmt.Errorf("amount cannot be negative"))
	})

	t.Run("✅ should update an expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		// When
		expense.Id = 1
		expense.Description = "Dinner"
		expenses.Update(expense)

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
		expenses.Add(expense)

		// When
		expense.Id = 2
		expense.Description = "Dinner"
		expenses.Update(expense)

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
		expenses.Add(expense)

		// When
		expenses.Delete(1)

		// Then
		asserts.Equal(0, len(expenses))
	})

	t.Run("❌ should not delete an non-existent expense", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		// When
		expenses.Delete(2)

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
		expenses.Add(expense)

		expectedMessage := dsl.JoinMessage(expenses)
		expectedMessage = fmt.Sprintf(models.HeaderFormat + "\n" + expectedMessage)

		// When
		result := dsl.OutputToString(expenses.List)

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
		result := dsl.OutputToString(expenses.List)

		// Then
		asserts.Equal(models.HeaderFormat+"\n", result)
	})

	t.Run("✅ should print the summary of all expenses", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 15
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		// When
		result := dsl.OutputToString(expenses.Summary)

		// Then
		asserts.Equal("Total expenses: 35\n", result)
	})

	t.Run("✅ should print the summary of all expenses for a specific month of the current year", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 15
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 50
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		expenses[0].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[1].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[2].CreatedAt = time.Date(time.Now().Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC)

		// When
		result := dsl.OutputToString(func() {
			expenses.SummaryForMonth(time.January)
		})

		// Then
		asserts.Equal("Total expenses: 35\n", result)
	})

	t.Run("✅ should print 0 when there are no expenses for a specific month of the current year", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 15
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 50
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		expenses[0].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[1].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[2].CreatedAt = time.Date(time.Now().Year()-1, time.February, 1, 0, 0, 0, 0, time.UTC)

		// When
		result := dsl.OutputToString(func() {
			expenses.SummaryForMonth(time.February)
		})

		// Then
		asserts.Equal("Total expenses: 0\n", result)
	})

	t.Run("✅ should print the expenses from a specific month and current year taking on account the updated at", func(t *testing.T) {
		// Given
		amount := 20
		description := "Lunch"
		expenses := models.Expenses{}

		expense := models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 15
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 50
		description = "Dinner"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		amount = 100
		description = "Lunch"
		expense = models.Expense{Amount: amount, Description: description}
		expenses.Add(expense)

		expenses[0].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[1].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[2].CreatedAt = time.Date(time.Now().Year()-1, time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[3].CreatedAt = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.UTC)

		updateAt := time.Date(time.Now().Year(), time.January, 25, 0, 0, 0, 0, time.UTC)
		expenses[0].UpdatedAt = &updateAt

		updateAt2 := time.Date(time.Now().Year(), time.February, 1, 0, 0, 0, 0, time.UTC)
		expenses[1].UpdatedAt = &updateAt2

		updateAt3 := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
		expenses[3].UpdatedAt = &updateAt3

		// When
		result := dsl.OutputToString(func() {
			expenses.SummaryForMonth(time.January)
		})

		// Then
		asserts.Equal("Total expenses: 20\n", result)
	})
}
