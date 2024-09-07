package main

import (
	"time"

	monthview "github.com/anotherhadi/calendar/month-view"
	calendar "github.com/anotherhadi/markdown-calendar"
	purple "github.com/anotherhadi/purple-apps"
)

type model struct {
	Calendars     []*calendar.Calendar
	Width, Height int

	CurrentDay, CurrentMonth, CurrentYear int
	FocusDay, FocusMonth, FocusYear       *int

	CurrentView string // "year", "month", "week", "day", "event", "new-event", "search", "edit-event"

	MonthModel monthview.Model
}

func initModel() model {
	m := model{
		Width:  0,
		Height: 0,
	}

	calendars := []*calendar.Calendar{}
	for _, p := range purple.Config.Calendar.Paths {
		cal, err := calendar.Read(p)
		if err != nil {
			continue
		}
		calendars = append(calendars, &cal)
	}
	m.Calendars = calendars

	m.CurrentView = purple.Config.Calendar.DefaultView

	t := time.Now()
	m.CurrentDay = t.Day()
	m.CurrentMonth = int(t.Month())
	m.CurrentYear = t.Year()
	focusDay := m.CurrentDay
	focusMonth := m.CurrentMonth
	focusYear := m.CurrentYear
	m.FocusDay = &focusDay
	m.FocusMonth = &focusMonth
	m.FocusYear = &focusYear

	m.MonthModel = monthview.Model{
		CurrentDay:   m.CurrentDay,
		CurrentMonth: m.CurrentMonth,
		CurrentYear:  m.CurrentYear,
		FocusDay:     m.FocusDay,
		FocusMonth:   m.FocusMonth,
		FocusYear:    m.FocusYear,
		Width:        m.Width,
		Height:       m.Height,
		Calendars:    calendars,
	}

	return m
}
