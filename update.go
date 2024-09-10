package main

import (
	neweventview "github.com/anotherhadi/calendar/new_event_view"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.MonthModel, _ = m.MonthModel.Update(msg)
		m.NewEventModel, _ = m.NewEventModel.Update(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if *m.IsNewEventView {
				*m.IsNewEventView = false
				return m, nil
			} else {
				return m, tea.Quit
			}
		case "n": // New event
			*m.IsNewEventView = !*m.IsNewEventView
			if *m.IsNewEventView {
				m.NewEventModel = neweventview.NewModel(m.Calendars, m.IsNewEventView, *m.FocusDay, *m.FocusMonth, *m.FocusYear)
				m.NewEventModel, _ = m.NewEventModel.Update(tea.WindowSizeMsg{Width: m.Width, Height: m.Height})
				return m, tea.Batch(m.NewEventModel.Init())
			}
		case "t": // Go to today
			*m.FocusDay = m.CurrentDay
			*m.FocusMonth = m.CurrentMonth
			*m.FocusYear = m.CurrentYear
			return m, nil

		default:
			if *m.IsNewEventView {
				m.NewEventModel, cmd = m.NewEventModel.Update(msg)
				return m, cmd
			}
			if m.CurrentView == "month" {
				m.MonthModel, cmd = m.MonthModel.Update(message)
			}
		}

	default:
		if *m.IsNewEventView {
			m.NewEventModel, cmd = m.NewEventModel.Update(message)
			return m, cmd
		}
		if m.CurrentView == "month" {
			m.MonthModel, cmd = m.MonthModel.Update(message)
		}
	}

	return m, cmd
}
