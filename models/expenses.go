package models

import (
	"fmt"
	"time"
)

// TODO: add type Amount int -> for validating amount

type (
	Expense struct {
		Id          int
		Amount      int
		Description string
		CreatedAt   time.Time
		UpdatedAt   *time.Time
	}

	Expenses []*Expense
)

const (
	HeaderFormat         = "|ID|Description|Amount|Created At|Updated At|"
	ExpensesStringFormat = "|%-3d|%-10s|%-6d|%-10s|%-10s|"
	DateFormat           = time.DateOnly
)

func (e *Expense) Print() {
	if e.UpdatedAt == nil {
		fmt.Printf(ExpensesStringFormat, e.Id, e.Description, e.Amount, e.CreatedAt.Format(time.DateOnly), "")
		return
	}
	fmt.Printf(ExpensesStringFormat, e.Id, e.Description, e.Amount, e.CreatedAt.Format(DateFormat), e.UpdatedAt.Format(DateFormat))
}
