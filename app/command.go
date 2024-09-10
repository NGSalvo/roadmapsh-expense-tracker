package app

import (
	"expense-tracker/models"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type (
	CommandLine interface {
		Run()
	}

	commandLine struct {
		Store models.Store
	}
)

func NewCommandLine(store models.Store) CommandLine {
	return &commandLine{store}
}

func (c *commandLine) Run() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s <command> [arguments]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  add       Add an expense\n")
		fmt.Fprintf(os.Stderr, "  list      List expenses\n")
		fmt.Fprintf(os.Stderr, "  delete    Delete expenses\n")
		fmt.Fprintf(os.Stderr, "  summary   Summary expenses\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	switch flag.Arg(0) {
	case "add":
		c.addExpenseCommand()
	case "list":
		c.listExpensesCommand()
	case "delete":
		c.deleteExpensesCommand()
	case "summary":
		c.summaryExpensesCommand()
	default:
		flag.Usage()
		os.Exit(1)
	}

}

func (c *commandLine) addExpenseCommand() {
	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCommand.String("description", "", "Description of the expense")
	amount := addCommand.Int("amount", 0, "Amount of the expense")
	addCommand.Parse(os.Args[2:])

	if *description == "" || *amount == 0 {
		log.Fatal("Description and amount are required")
	}

	error := c.Store.Add(models.Expense{Description: *description, Amount: *amount})

	if error != nil {
		log.Fatal(error)
	}
}

func (c *commandLine) listExpensesCommand() {
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	listCommand.Parse(os.Args[2:])

	c.Store.List()
}

func (c *commandLine) deleteExpensesCommand() {
	deleteCommand := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteId := deleteCommand.Int("id", 0, "ID of the expense")

	deleteCommand.Parse(os.Args[2:])
	if *deleteId == 0 {
		log.Fatal("ID is required")
	}

	error := c.Store.Delete(*deleteId)

	if error != nil {
		log.Fatal(error)
	}
}

func (c *commandLine) summaryExpensesCommand() {
	summaryCommand := flag.NewFlagSet("summary", flag.ExitOnError)
	summaryMonth := summaryCommand.Int("month", 0, "Month of the summary")
	summaryCommand.Parse(os.Args[2:])
	if *summaryMonth == 0 {
		c.Store.Summary()
		os.Exit(0)
	}
	c.Store.SummaryForMonth(time.Month(*summaryMonth))
}
