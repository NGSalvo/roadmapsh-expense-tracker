package models

import "time"

type (
	Store interface {
		Add(expense Expense) error
		List()
		Update(expense Expense) error
		Delete(id int)
		Summary()
		SummaryForMonth(month time.Month)
	}
)
