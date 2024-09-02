package main

import "time"

func main() {

}

type (
	Expense struct {
		Amount      int
		Description string
		CreatedAt   *time.Time
		UpdatedAt   *time.Time
	}

	Expenses []*Expense
)

func (e *Expenses) AddExpense(expense Expense) {
	*e = append(*e, &expense)
}
