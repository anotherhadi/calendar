package main

import (
	month "github.com/anotherhadi/calendar/month_view"
	week "github.com/anotherhadi/calendar/week_view"
)

func (m model) View() string {
	var s string
	var help string
	if m.CurrentView == "month" {
		s = m.MonthModel.View()
		help = m.MonthModel.Help.View(month.Keys)
	}
	if m.CurrentView == "week" {
		s = m.WeekModel.View()
		help = m.WeekModel.Help.View(week.Keys)
	}
	if m.CurrentView == "new_event" {
		s = m.NewEventModel.View()
	}

	return m.drawHelp(s, help)
}
