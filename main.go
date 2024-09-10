package main

import (
	"expense-tracker/app"
	"expense-tracker/stores"
)

func main() {
	app.NewCommandLine(stores.NewCsvStore("test.csv")).Run()
}
