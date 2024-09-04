package models

import (
	"fmt"
	"time"
)

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

func (e *Expenses) Add(expense Expense) error {
	expense.Id = e.assignId()
	expense.CreatedAt = time.Now()

	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	*e = append(*e, &expense)
	return nil
}

func (e *Expenses) assignId() int {
	currentMaxId := 0
	for _, expense := range *e {
		if expense.Id > currentMaxId {
			currentMaxId = expense.Id
		}
	}
	return (currentMaxId + 1)
}

func (e *Expenses) Update(expense Expense) error {
	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	for _, item := range *e {
		if item.Id == expense.Id {
			item.Amount = expense.Amount
			item.Description = expense.Description
			updatedAt := time.Now()
			item.UpdatedAt = &updatedAt
			return nil
		}
	}
	return fmt.Errorf("expense with ID %d not found", expense.Id)
}

func (e *Expenses) Delete(id int) {
	for i, item := range *e {
		if item.Id == id {
			*e = append((*e)[:i], (*e)[i+1:]...)
			break
		}
	}
}

func (e *Expenses) List() {
	fmt.Printf(HeaderFormat + "\n")
	for _, expense := range *e {
		expense.Print()
	}
}

func (e *Expenses) Summary() {
	total := 0
	for _, expense := range *e {
		total += expense.Amount
	}
	fmt.Printf("Total expenses: %d\n", total)
}

func (e *Expenses) SummaryForMonth(month time.Month) {
	total := 0
	for _, expense := range *e {
		if expense.UpdatedAt == nil {
			if e.isValidSummaryInput(&expense.CreatedAt, month) {
				total += expense.Amount
			}
			continue
		}

		if e.isValidSummaryInput(expense.UpdatedAt, month) {
			total += expense.Amount
		}
	}
	fmt.Printf("Total expenses: %d\n", total)
}

func (e *Expenses) isValidSummaryInput(date *time.Time, month time.Month) bool {
	isCurrentYear := date.Year() == time.Now().Year()
	isSameMonth := date.Month() == month

	if isSameMonth && isCurrentYear {
		return true
	}
	return false
}
