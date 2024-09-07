package main

import (
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var str string
	if m.CurrentView == "month" {
		str = m.MonthModel.View()
	}

	if m.IsNewEventView {
		str = lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, m.NewEventModel.View(), lipgloss.WithWhitespaceChars("您好"), lipgloss.WithWhitespaceForeground(purple.Colors.Muted))
	}
	return str
}
