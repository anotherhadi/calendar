package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.MonthModel, _ = m.MonthModel.Update(msg)
		msg.Height -= 6
		msg.Width -= 10
		m.NewEventModel, _ = m.NewEventModel.Update(msg)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.IsNewEventView {
				m.IsNewEventView = false
				return m, nil
			} else {
				return m, tea.Quit
			}
		case "n":
			m.IsNewEventView = !m.IsNewEventView
			if m.IsNewEventView {
				return m, tea.Batch(m.NewEventModel.Init())
			}
		default:
			if m.IsNewEventView {
				m.NewEventModel, cmd = m.NewEventModel.Update(msg)
				return m, cmd
			}
			if m.CurrentView == "month" {
				m.MonthModel, cmd = m.MonthModel.Update(message)
			}
		}

	default:
		if m.IsNewEventView {
			m.NewEventModel, cmd = m.NewEventModel.Update(message)
			return m, cmd
		}
		if m.CurrentView == "month" {
			m.MonthModel, cmd = m.MonthModel.Update(message)
		}
	}

	return m, cmd
}
