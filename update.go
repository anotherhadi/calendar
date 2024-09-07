package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.MonthModel, _ = m.MonthModel.Update(message)

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		default:
			if m.CurrentView == "month" {
				m.MonthModel, _ = m.MonthModel.Update(message)
			}
		}
	}

	return m, nil
}
