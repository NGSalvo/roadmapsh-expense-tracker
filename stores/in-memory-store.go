package stores

import (
	"expense-tracker/models"
	"fmt"
	"time"
)

type InMemoryStore struct {
	Expenses *models.Expenses
}

func NewInMemoryStore() models.Store {
	return &InMemoryStore{
		Expenses: &models.Expenses{},
	}
}

func (s *InMemoryStore) Add(expense models.Expense) error {
	expense.Id = s.assignId()
	expense.CreatedAt = time.Now()

	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	*s.Expenses = append(*s.Expenses, &expense)

	return nil
}

func (s *InMemoryStore) assignId() int {
	currentMaxId := 0
	for _, expense := range *s.Expenses {
		if expense.Id > currentMaxId {
			currentMaxId = expense.Id
		}
	}
	return (currentMaxId + 1)
}

func (s *InMemoryStore) Update(expense models.Expense) error {
	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	for _, item := range *s.Expenses {
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

func (s *InMemoryStore) Delete(id int) error {
	foundItem := false
	for i, item := range *s.Expenses {
		if item.Id == id {
			*s.Expenses = append((*s.Expenses)[:i], (*s.Expenses)[i+1:]...)
			break
		}
	}

	if !foundItem {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return nil
}

func (s *InMemoryStore) List() {
	fmt.Printf(models.HeaderFormat + "\n")
	for _, expense := range *s.Expenses {
		expense.Print()
	}
}

func (s *InMemoryStore) Summary() {
	total := 0
	for _, expense := range *s.Expenses {
		total += expense.Amount
	}
	fmt.Printf("Total expenses: %d\n", total)
}

func (s *InMemoryStore) SummaryForMonth(month time.Month) {
	total := 0
	for _, expense := range *s.Expenses {
		if expense.UpdatedAt == nil {
			if s.isValidSummaryInput(&expense.CreatedAt, month) {
				total += expense.Amount
			}
			continue
		}

		if s.isValidSummaryInput(expense.UpdatedAt, month) {
			total += expense.Amount
		}
	}
	fmt.Printf("Total expenses: %d\n", total)
}

func (s *InMemoryStore) isValidSummaryInput(date *time.Time, month time.Month) bool {
	isCurrentYear := date.Year() == time.Now().Year()
	isSameMonth := date.Month() == month

	if isSameMonth && isCurrentYear {
		return true
	}
	return false
}
