package main

import (
	month "github.com/anotherhadi/calendar/month_view"
)

func (m model) View() string {
	var s string
	var help string
	if m.CurrentView == "month" {
		s = m.MonthModel.View()
		help = m.MonthModel.Help.View(month.Keys)
	}
	if m.CurrentView == "new_event" {
		s = m.NewEventModel.View()
	}

	return m.drawHelp(s, help)
}
