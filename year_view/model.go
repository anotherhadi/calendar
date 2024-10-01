package year

import (
	"github.com/anotherhadi/calendar/style"
	"github.com/charmbracelet/bubbles/v2/help"
)

type Model struct {
	currentDay, currentMonth, currentYear int
	focusDay, focusMonth, focusYear       *int

	keys          keyMap
	Help          help.Model
	width, height int
}

func NewModel(
	currentDay, currentMonth, currentYear int,
	focusDay, focusMonth, focusYear *int,
) Model {
	help := help.New()
	help.Styles = style.GetHelpStyles()

	m := Model{
		currentDay:   currentDay,
		currentMonth: currentMonth,
		currentYear:  currentYear,
		focusDay:     focusDay,
		focusMonth:   focusMonth,
		focusYear:    focusYear,

		keys:   Keys,
		Help:   help,
		width:  0,
		height: 0,
	}

	return m
}
