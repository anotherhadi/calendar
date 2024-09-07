package monthview

import (
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) getNumberOfEvents(year, month, day int) int {
	n := 0
	for _, cal := range m.Calendars {
		n += len(cal.GetEventsByDate(year, month, day))
	}
	return n
}

func (m Model) getEvents(year, month, day int) []calendar.Event {
	n := []calendar.Event{}
	for _, cal := range m.Calendars {
		n = append(n, cal.GetEventsByDate(year, month, day)...)
	}
	return n
}

func (m Model) getNumberOfEventsInMonth(year, month int) int {
	n := 0
	for _, cal := range m.Calendars {
		n += len(cal.GetEventsByMonth(year, month))
	}
	return n
}

// getHeaders returns the headers for the month view depending on the width of the terminal
func getHeaders(width int) []string {
	var headers []string
	if width < 35 {
		headers = []string{"M", "T", "W", "T", "F", "S", "S"}
	} else if width < 50 {
		headers = []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	} else if width < 70 {
		headers = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	} else {
		headers = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	}
	return headers
}

func today(str string, hover bool) string {
	s := lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	// if hover {
	// 	border = border.Background(purple.Colors.Muted)
	// }
	// inside := lipgloss.NewStyle().Background(purple.Colors.Accent).Foreground(purple.GetFgColor(purple.Colors.Accent))
	//
	// var result string
	// result += border.Render("")
	// result += inside.Render(str)
	// result += border.Render("")
	return s.Render("󰃭", str)
}
