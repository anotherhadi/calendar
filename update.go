package main

import (
	newevent "github.com/anotherhadi/calendar/new_event_view"
	"github.com/anotherhadi/calendar/utils"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func (m model) UpdateSize(msg tea.WindowSizeMsg) (model, tea.Cmd) {
	m.Width = msg.Width
	m.Height = msg.Height
	m.MonthModel, _ = m.MonthModel.Update(msg)
	m.NewEventModel, _ = m.NewEventModel.Update(msg)
	return m, nil
}

func (m model) UpdateFocusedView(message tea.Msg) (model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.CurrentView {
	case "month":
		m.MonthModel, cmd = m.MonthModel.Update(message)
	case "new_event":
		m.NewEventModel, cmd = m.NewEventModel.Update(message)
	}

	return m, cmd
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		return m.UpdateSize(msg)

	case utils.ChangeFocusViewMsg:
		if msg.View == "new_event" {
			m.NewEventModel = newevent.NewModel(m.Calendars, m.View(), m.CurrentView)
		}
		m.CurrentView = msg.View
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m.UpdateFocusedView(message)
}
