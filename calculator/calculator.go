package calculator

// https://godoc.org/github.com/rivo/tview

import (
	"errors"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
	"strings"
)

var valid_numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var valid_operands = []string{"+", "-", "*", "/"}
var valid_grouping = []string{"(", ")"}

type Calculator struct {
	*tview.Flex
	app        *tview.Application
	resultView *tview.TextView
	inputBox   *tview.InputField
	operation  string
	helpView   *tview.TextView
	compute    *Compute
}

func NewCalculator(app *tview.Application) *Calculator {
	var cal = Calculator{
		Flex:    tview.NewFlex(),
		app:     app,
		compute: NewCompute(),
	}
	cal.init()
	return &cal
}

func (c *Calculator) init() {
	c.Flex.SetDirection(tview.FlexRow)

	// Result View
	c.resultView = tview.NewTextView()
	c.resultView.SetBorder(true).SetTitle("  Results  ")

	// Input Box
	c.inputBox = tview.NewInputField().
		SetPlaceholder("e.g. 12 + 5").
		SetDoneFunc(c.done).
		SetAcceptanceFunc(c.validate)

	c.inputBox.SetBorderPadding(2, 2, 5, 5)

	// Help View
	c.helpView = tview.NewTextView()
	c.helpView.SetBorder(true).SetTitle("   help   ")
	c.populateHelp()

	c.Flex.AddItem(c.resultView, 0, 2, false).
		AddItem(c.inputBox, 0, 1, true).
		AddItem(c.helpView, 0, 2, true)
}

func (c *Calculator) done(key tcell.Key) {
	if key == tcell.KeyEscape {
		c.app.Stop()
	}

	if key == tcell.KeyTab {
		c.inputBox.SetText("")
	}

	if key == tcell.KeyEnter {
		c.calculate()
	}
}

func (c *Calculator) validate(textToCheck string, lastChar rune) bool {
	theChar := string(lastChar)

	for _, i := range valid_numbers {
		if i == theChar {
			return true
		}
	}

	for _, i := range valid_grouping {
		if i == theChar {
			return true
		}
	}

	if len(textToCheck) == 1 && theChar != "-" {
		return false
	}

	for _, i := range valid_operands {
		if i == theChar {
			return true
		}
	}

	return false
}

func (c *Calculator) calculate() {
	c.operation = c.inputBox.GetText()
	err := c.sanitize()
	if err != nil {
		c.resultView.SetText(err.Error())
		return
	}
	output, err := c.compute.process(c.operation)
	if err != nil {
		c.resultView.SetText(err.Error())
		return
	}

	s := strconv.FormatFloat(output, 'E', -1, 64) // todo: format this correctly
	c.resultView.SetText(s)
}

func (c *Calculator) sanitize() error {
	c.operation = strings.Trim(c.operation, "\n \t")
	if len(c.operation) == 0 {
		return errors.New("Nothing to do!")
	}
	lastChar := string(c.operation[len(c.operation)-1])
	isOperator := false
	for _, i := range valid_operands {
		if lastChar == i {
			isOperator = true
			break
		}
	}

	if isOperator == true {
		return errors.New("Invalid input!")
	}

	return nil
}

func (c *Calculator) populateHelp() {
	c.helpView.SetTextColor(tcell.ColorYellow)
	c.helpView.SetText(`
      tab   - clear input
      esc   - to exit
  `)
}
