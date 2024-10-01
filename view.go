package main

import (
	month "github.com/anotherhadi/calendar/month_view"
	week "github.com/anotherhadi/calendar/week_view"
	year "github.com/anotherhadi/calendar/year_view"
)

func (m model) View() string {
	var s string
	var help string
	switch m.CurrentView {
	case "month":
		s = m.MonthModel.View()
		help = m.MonthModel.Help.View(month.Keys)
	case "week":
		s = m.WeekModel.View()
		help = m.WeekModel.Help.View(week.Keys)
	case "year":
		s = m.YearModel.View()
		help = m.YearModel.Help.View(year.Keys)
	case "new_event":
		s = m.NewEventModel.View()
	}

	return m.drawHelp(s, help)
}
