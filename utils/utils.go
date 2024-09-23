package utils

import (
	"regexp"

	calendar "github.com/anotherhadi/markdown-calendar"
	oldTea "github.com/charmbracelet/bubbletea"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func RemoveAnsiStyle(s string) string {
	return regexp.MustCompile(`\x1b[[\d;]*m`).ReplaceAllString(s, "")
}

func PtrCalendarsToCalendars(ptrCalendars []*calendar.Calendar) []calendar.Calendar {
	calendars := make([]calendar.Calendar, len(ptrCalendars))
	for i, ptrCal := range ptrCalendars {
		calendars[i] = *ptrCal
	}
	return calendars
}

func TruncateString(s string, n int) string {
	if len(s) <= n {
		return s
	}
	if n <= 3 {
		return ""
	}
	return s[:n-3] + "..."
}

func WrapOldBubbleteaCmd(oldCmd oldTea.Cmd) (newCmd tea.Cmd) {
	newCmd = func() tea.Msg {
		if oldCmd != nil {
			return oldCmd()
		}
		return nil
	}
	return newCmd
}
