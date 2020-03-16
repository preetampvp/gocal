package main

import (
	"github.com/preetampvp/gocal/calculator"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	calculator := calculator.NewCalculator()
	if err := app.SetRoot(calculator, true).Run(); err != nil {
		panic(err)
	}
}
