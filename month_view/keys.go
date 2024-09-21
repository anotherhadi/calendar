package month

import "github.com/charmbracelet/bubbles/v2/key"

type keyMap struct {
	PreviousDay   key.Binding
	NextDay       key.Binding
	PreviousWeek  key.Binding
	NextWeek      key.Binding
	PreviousMonth key.Binding
	NextMonth     key.Binding
	PreviousYear  key.Binding
	NextYear      key.Binding
	NewEvent      key.Binding
	DayView       key.Binding
	WeekView      key.Binding
	YearView      key.Binding
	Today         key.Binding

	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PreviousDay, k.PreviousMonth, k.NewEvent, k.Today},
		{k.NextWeek, k.NextYear, k.DayView, k.Quit},
		{k.PreviousWeek, k.PreviousYear, k.WeekView, k.Help},
		{k.NextDay, k.NextMonth, k.YearView},
	}
}

var Keys = keyMap{
	PreviousWeek: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "previous week"),
	),
	NextWeek: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "next week"),
	),
	PreviousDay: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "previous day"),
	),
	NextDay: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "next day"),
	),
	PreviousMonth: key.NewBinding(
		key.WithKeys("shift+h"),
		key.WithHelp("H", "previous month"),
	),
	NextMonth: key.NewBinding(
		key.WithKeys("shift+l"),
		key.WithHelp("L", "next month"),
	),
	PreviousYear: key.NewBinding(
		key.WithKeys("shift+k"),
		key.WithHelp("K", "previous year"),
	),
	NextYear: key.NewBinding(
		key.WithKeys("shift+j"),
		key.WithHelp("J", "next year"),
	),
	NewEvent: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "new event"),
	),
	DayView: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "day view"),
	),
	WeekView: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "week view"),
	),
	YearView: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "year view"),
	),
	Today: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "today"),
	),

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
