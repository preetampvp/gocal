package calculator

import (
	"github.com/rivo/tview"
)

type Calculator struct {
	*tview.Flex
	resultBox *tview.TextView
	inputBox  *tview.InputField
}

func NewCalculator() *Calculator {
	var cal = Calculator{
		Flex: tview.NewFlex(),
	}
	cal.init()
	return &cal
}

func (c *Calculator) init() {
	c.Flex.SetDirection(tview.FlexRow)

	c.resultBox = tview.NewTextView()
	c.resultBox.SetBorder(true).SetTitle("Results")

	c.inputBox = tview.NewInputField().
		SetPlaceholder("e.g. 12 + 5").
		SetChangedFunc(c.inputChange)

	c.Flex.AddItem(c.resultBox, 0, 3, false).
		AddItem(c.inputBox, 0, 1, true)
}

func (c *Calculator) inputChange(text string) {
	c.resultBox.SetText(text)
}
