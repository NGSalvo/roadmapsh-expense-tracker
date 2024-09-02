package main

import "time"

func main() {

}

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

func (e *Expenses) AddExpense(expense Expense) {
	expense.Id = e.assignId()
	expense.CreatedAt = time.Now()
	*e = append(*e, &expense)
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

func (e *Expenses) UpdateExpense(expense Expense) {
	for _, item := range *e {
		if item.Id == expense.Id {
			item.Amount = expense.Amount
			item.Description = expense.Description
			updatedAt := time.Now()
			item.UpdatedAt = &updatedAt
			break
		}
	}
}
