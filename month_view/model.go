package month

import (
	"github.com/anotherhadi/calendar/style"
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/charmbracelet/bubbles/v2/help"
)

type Model struct {
	currentDay, currentMonth, currentYear int
	focusDay, focusMonth, focusYear       *int

	calendar    *calendar.Calendar
	rawCalendar *[][]string

	keys          keyMap
	Help          help.Model
	width, height int

	msg *string
}

func NewModel(
	currentDay, currentMonth, currentYear int,
	focusDay, focusMonth, focusYear *int,
	calendar *calendar.Calendar,
) Model {
	help := help.New()
	help.Styles = style.GetHelpStyles()

	testmsg := "test"

	m := Model{
		currentDay:   currentDay,
		currentMonth: currentMonth,
		currentYear:  currentYear,
		focusDay:     focusDay,
		focusMonth:   focusMonth,
		focusYear:    focusYear,
		calendar:     calendar,
		rawCalendar:  &[][]string{},
		msg:          &testmsg,

		keys:   Keys,
		Help:   help,
		width:  0,
		height: 0,
	}

	return m
}
