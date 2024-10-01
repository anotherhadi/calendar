package year

import "github.com/charmbracelet/bubbles/v2/key"

type keyMap struct {
	PreviousMonth key.Binding
	NextMonth     key.Binding
	PreviousYear  key.Binding
	NextYear      key.Binding
	NewEvent      key.Binding
	DayView       key.Binding
	WeekView      key.Binding
	MonthView     key.Binding
	Today         key.Binding

	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PreviousMonth, k.NewEvent, k.Today},
		{k.NextYear, k.DayView, k.Quit},
		{k.PreviousYear, k.WeekView, k.Help},
		{k.NextMonth, k.MonthView},
	}
}

var Keys = keyMap{
	PreviousMonth: key.NewBinding(
		key.WithKeys("left", "h", "shift+h"),
		key.WithHelp("←/h", "previous month"),
	),
	NextMonth: key.NewBinding(
		key.WithKeys("right", "l", "shift+l"),
		key.WithHelp("→/l", "next month"),
	),
	PreviousYear: key.NewBinding(
		key.WithKeys("up", "k", "shift+k"),
		key.WithHelp("↑/k", "previous year"),
	),
	NextYear: key.NewBinding(
		key.WithKeys("down", "j", "shift+j"),
		key.WithHelp("↓/j", "next year"),
	),
	NewEvent: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "new event"),
	),
	DayView: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "day view"),
	),
	WeekView: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "week view"),
	),
	MonthView: key.NewBinding(
		key.WithKeys("m", "enter"),
		key.WithHelp("m/󱞥", "month view"),
	),
	Today: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp("t", "today"),
	),

	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
