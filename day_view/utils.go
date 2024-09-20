package month

import (
	"strconv"

	"github.com/anotherhadi/calendar/style"
	calendar "github.com/anotherhadi/markdown-calendar"
)

func ptrCalendarsToCalendars(ptrCalendars []*calendar.Calendar) []calendar.Calendar {
	var calendars []calendar.Calendar
	for _, c := range ptrCalendars {
		calendars = append(calendars, *c)
	}
	return calendars
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

// Place the previous month's days in the calendar
func (m Model) mutedPreviousDays(rows [][]string) [][]string {
	tmp := []string{}
	for j := 0; j < calendar.DayOfWeek(1, *m.focusMonth, *m.focusYear); j++ {
		tmp = append(
			tmp,
			style.OutsideCellStyle.Render(
				strconv.Itoa(
					calendar.DaysInMonth(
						*m.focusMonth-1,
						*m.focusYear,
					)-calendar.DayOfWeek(
						1,
						*m.focusMonth,
						*m.focusYear,
					)+j+1,
				),
			),
		)
	}
	rows = append(rows, tmp)
	return rows
}

// Place the next month's days in the calendar
func (m Model) mutedNextDays(rows [][]string) [][]string {
	// Place the next month's days in the calendar
	for i := 1; len(rows[len(rows)-1]) < 7; i++ {
		rows[len(rows)-1] = append(
			rows[len(rows)-1],
			style.OutsideCellStyle.Render(strconv.Itoa(i)),
		)
	}
	return rows
}
