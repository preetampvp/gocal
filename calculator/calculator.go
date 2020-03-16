package calculator

// https://godoc.org/github.com/rivo/tview

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var valid_chars = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "+", "-", "*", "/", "(", ")"}

type Calculator struct {
	*tview.Flex
	app        *tview.Application
	resultView *tview.TextView
	inputBox   *tview.InputField
	helpView   *tview.TextView
}

func NewCalculator(app *tview.Application) *Calculator {
	var cal = Calculator{
		Flex: tview.NewFlex(),
		app:  app,
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
		SetChangedFunc(c.inputChange).
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

func (c *Calculator) inputChange(text string) {
	// c.resultView.SetText(text)
}

func (c *Calculator) done(key tcell.Key) {
	if key == tcell.KeyEscape {
		c.app.Stop()
	}
	if key == tcell.KeyTab {
		c.inputBox.SetText("")
	}
}

func (c *Calculator) validate(textToCheck string, lastChar rune) bool {
	theChar := string(lastChar)

	for _, i := range valid_chars {
		if i == theChar {
			return true
		}
	}

	return false
}

func (c *Calculator) populateHelp() {
	c.helpView.SetTextColor(tcell.ColorYellow)
	c.helpView.SetText(`
      tab   - clear input
      esc   - to exit
  `)
}
