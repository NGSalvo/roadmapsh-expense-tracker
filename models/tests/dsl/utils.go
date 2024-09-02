package dsl

import (
	"bytes"
	"expense-tracker/models"
	"fmt"
	"io"
	"os"
	"strings"
)

func OutputToString(callback func()) string {

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()

	// Save the original stdout
	oldStdout := os.Stdout

	// Assign the write end of the pipe to stdout
	os.Stdout = w

	// Ensure that stdout is restored after the test
	defer func() {
		os.Stdout = oldStdout
		w.Close()
	}()

	// Call the function or code that prints to stdout
	callback()

	// Close the write end of the pipe to signal EOF
	w.Close()

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func JoinMessage(expenses models.Expenses) string {
	message := []string{}
	for _, expense := range expenses {
		if expense.UpdatedAt == nil {
			message = append(message, fmt.Sprintf(models.ExpensesStringFormat, expense.Id, expense.Description, expense.Amount, expense.CreatedAt.Format(models.DateFormat), ""))
			continue
		}
		message = append(message, fmt.Sprintf(models.ExpensesStringFormat, expense.Id, expense.Description, expense.Amount, expense.CreatedAt.Format(models.DateFormat), expense.UpdatedAt.Format(models.DateFormat)))
	}
	return strings.Join(message, "\n")
}
