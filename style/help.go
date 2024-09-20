package style

import (
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/bubbles/v2/help"
	"github.com/charmbracelet/lipgloss"
)

func GetHelpStyles() help.Styles {
	keyStyle := lipgloss.NewStyle().Foreground(purple.Colors.LightGray)

	descStyle := lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	sepStyle := lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	return help.Styles{
		ShortKey:       keyStyle,
		ShortDesc:      descStyle,
		ShortSeparator: sepStyle,
		Ellipsis:       sepStyle,
		FullKey:        keyStyle,
		FullDesc:       descStyle,
		FullSeparator:  sepStyle,
	}
}
