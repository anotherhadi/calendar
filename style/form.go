package style

import (
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func GetFormTheme() *huh.Theme {
	theme := huh.ThemeBase()
	theme.Focused.Title = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Blurred.Title = lipgloss.NewStyle().Foreground(purple.Colors.LightGray)
	theme.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true).BorderForeground(purple.Colors.Accent)
	theme.Blurred.Description = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	theme.Focused.Description = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	theme.Focused.TextInput.Prompt = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Blurred.TextInput.Prompt = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	theme.Focused.SelectSelector = lipgloss.NewStyle().Foreground(purple.Colors.Muted).SetString("> ")
	theme.Focused.SelectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Focused.UnselectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	theme.Blurred.SelectSelector = lipgloss.NewStyle().Foreground(purple.Colors.Muted).SetString("> ")
	theme.Blurred.SelectedOption = lipgloss.NewStyle().Foreground(purple.Colors.LightGray)
	theme.Blurred.UnselectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	return theme
}
