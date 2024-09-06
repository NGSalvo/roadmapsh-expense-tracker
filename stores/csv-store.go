package stores

import (
	"encoding/csv"
	"expense-tracker/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type csvStore struct {
	Expenses *models.Expenses
	filename string
}

func NewCsvStore(filename string) models.Store {
	fileExistAndCreate(filename)

	return &csvStore{
		Expenses: &models.Expenses{},
		filename: filename,
	}
}

func (s *csvStore) Add(expense models.Expense) error {
	expense.Id = s.assignId()
	expense.CreatedAt = time.Now()

	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	*s.Expenses = append(*s.Expenses, &expense)
	err := s.save()

	if err != nil {
		return err
	}

	return nil
}

func (s *csvStore) assignId() int {
	currentMaxId := 0
	for _, expense := range *s.Expenses {
		if expense.Id > currentMaxId {
			currentMaxId = expense.Id
		}
	}
	return (currentMaxId + 1)
}

func (s *csvStore) Update(expense models.Expense) error {
	if expense.Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	for _, item := range *s.Expenses {
		if item.Id == expense.Id {
			item.Amount = expense.Amount
			item.Description = expense.Description
			updatedAt := time.Now()
			item.UpdatedAt = &updatedAt
			break
		}
	}
	err := s.save()

	if err != nil {
		return err
	}

	return fmt.Errorf("expense with ID %d not found", expense.Id)
}

func (s *csvStore) Delete(id int) error {
	foundItem := false
	for i, item := range *s.Expenses {
		if item.Id == id {
			foundItem = true
			*s.Expenses = append((*s.Expenses)[:i], (*s.Expenses)[i+1:]...)
			break
		}
	}

	if !foundItem {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	err := s.save()

	if err != nil {
		return err
	}

	return nil
}

func (s *csvStore) List() {
	fmt.Printf(models.HeaderFormat + "\n")
	for _, expense := range *s.Expenses {
		expense.Print()
	}
}

func (s *csvStore) Summary() {
	total := 0
	for _, expense := range *s.Expenses {
		total += expense.Amount
	}
	fmt.Printf("Total expenses: %d\n", total)
}

func (s *csvStore) SummaryForMonth(month time.Month) {
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

func (s *csvStore) isValidSummaryInput(date *time.Time, month time.Month) bool {
	isCurrentYear := date.Year() == time.Now().Year()
	isSameMonth := date.Month() == month

	if isSameMonth && isCurrentYear {
		return true
	}
	return false
}

func fileExistAndCreate(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)

		if err != nil {
			panic("Error creating file: " + err.Error())
		}
		defer closeFile(file)

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writeHeader(writer)
	}
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic("Error closing file: " + err.Error())
	}
}

func writeHeader(writer *csv.Writer) {
	headers := strings.Split(models.HeaderFormat, "|")
	headers = headers[1 : len(headers)-1]
	err := writer.Write(headers)
	if err != nil {
		panic("Error writing headers: " + err.Error())
	}
}

func (s *csvStore) load() error {
	file, err := os.OpenFile(s.filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return err
		}
		amount, err := strconv.Atoi(record[2])
		if err != nil {
			return err
		}
		createdAt, err := time.Parse(models.DateFormat, record[3])
		if err != nil {
			return err
		}
		var updatedAt time.Time
		if record[4] != "" {
			updatedAt, err = time.Parse(models.DateFormat, record[4])
			if err != nil {
				return err
			}
		}
		expense := models.Expense{
			Id:          id,
			Amount:      amount,
			Description: record[1],
			CreatedAt:   createdAt,
			UpdatedAt:   &updatedAt,
		}
		*s.Expenses = append(*s.Expenses, &expense)
	}
	return nil
}

func (s *csvStore) save() error {
	file, err := os.OpenFile(s.filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// os.file
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writeHeader(writer)

	var records [][]string
	for _, expense := range *s.Expenses {
		if expense.UpdatedAt == nil {
			records = append(records, []string{
				strconv.Itoa(expense.Id),
				expense.Description,
				strconv.Itoa(expense.Amount),
				expense.CreatedAt.Format(models.DateFormat),
				"",
			})
			continue
		}

		records = append(records, []string{
			strconv.Itoa(expense.Id),
			expense.Description,
			strconv.Itoa(expense.Amount),
			expense.CreatedAt.Format(models.DateFormat),
			expense.UpdatedAt.Format(models.DateFormat),
		})
	}
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return nil
}
