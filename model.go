package main

import (
	// neweventview "github.com/anotherhadi/calendar/new_event_view"

	month "github.com/anotherhadi/calendar/month_view"
	newevent "github.com/anotherhadi/calendar/new_event_view"
	"github.com/anotherhadi/calendar/utils"
	calendar "github.com/anotherhadi/markdown-calendar"
	purple "github.com/anotherhadi/purple-apps"
)

type model struct {
	Calendar      calendar.Calendar
	Width, Height int

	CurrentDay, CurrentMonth, CurrentYear int
	FocusDay, FocusMonth, FocusYear       *int

	CurrentView string // "year", "month", "week", "day", "new_event"

	MonthModel    month.Model
	NewEventModel newevent.Model
}

func initModel() model {
	m := model{
		Width:  0,
		Height: 0,
	}

	m.Calendar = calendar.MergeCalendars(
		utils.PtrCalendarsToCalendars(calendar.GetPurpleCalendars()),
	)
	m.CurrentView = purple.Config.Calendar.DefaultView
	m.CurrentDay, m.CurrentMonth, m.CurrentYear = calendar.Today()
	focusDay, focusMonth, focusYear := m.CurrentDay, m.CurrentMonth, m.CurrentYear
	m.FocusDay, m.FocusMonth, m.FocusYear = &focusDay, &focusMonth, &focusYear
	m.MonthModel = month.NewModel(
		m.CurrentDay,
		m.CurrentMonth,
		m.CurrentYear,
		m.FocusDay,
		m.FocusMonth,
		m.FocusYear,
		&m.Calendar,
	)
	m.NewEventModel = newevent.NewModel(&m.Calendar, "", "")

	return m
}
