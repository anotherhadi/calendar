package style

import (
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle = lipgloss.NewStyle().
			Foreground(purple.Colors.Accent).
			Align(lipgloss.Center).
			Bold(true)
	BorderStyle = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	// Cells
	CellStyle        = lipgloss.NewStyle().Padding(0, 1)
	CellStyleHover   = lipgloss.NewStyle().Padding(0, 1).Background(purple.Colors.Muted)
	OutsideCellStyle = CellStyle.Foreground(purple.Colors.Muted)

	// Event
	EventStyle = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	// Other
	Notice = lipgloss.NewStyle().Foreground(purple.Colors.Muted).Align(lipgloss.Center)
)
