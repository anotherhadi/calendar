package month

import (
	"fmt"
	"strconv"
	"time"

	"github.com/anotherhadi/calendar/style"
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (m Model) drawCalendar() string {
	var rows [][]string = [][]string{}
	var hoverRow, hoverCol, todayRow, todayCol int

	rows = m.mutedPreviousDays(rows)
	for i := 1; i <= calendar.DaysInMonth(*m.focusMonth, *m.focusYear); i++ {
		if i == *m.focusDay {
			hoverRow = len(rows)
			hoverCol = len(rows[len(rows)-1])
		}

		if i == m.currentDay && *m.focusMonth == m.currentMonth && *m.focusYear == m.currentYear {
			todayRow = len(rows)
			todayCol = len(rows[len(rows)-1])
		}

		rows[len(rows)-1] = append(rows[len(rows)-1], strconv.Itoa(i))
		if calendar.DayOfWeek(i, *m.focusMonth, *m.focusYear) == 6 {
			rows = append(rows, []string{})
		}
	}
	if len(rows[len(rows)-1]) == 0 { // Delete the last row if it's empty
		rows = rows[:len(rows)-1]
	}
	rows = m.mutedNextDays(rows)

	cellWidth := int((m.width-1-7)/7) - 2
	cellHeight := int(
		(m.height - 1 - 1 - 3 - len(rows)) / len(rows),
	) // 1 for title, 1 for the notice, 3 for header

	if cellHeight > 1 {
		for row := range rows {
			for col := range row {
				day, _ := strconv.Atoi(rows[row][col])

				nevents := calendar.GetNumberOfEventsByDate(
					*m.focusYear,
					*m.focusMonth,
					day,
					ptrCalendarsToCalendars(m.calendars),
				)
				rows[row][col] += "\n"

				if nevents < 1 {
					continue
				}

				var s lipgloss.Style

				s = style.EventStyle
				if row+1 == hoverRow && col == hoverCol {
					s = s.UnsetForeground().Foreground(purple.Colors.Accent)
				}
				var e string
				if cellWidth >= 11 {
					e = "event"
					if nevents > 1 {
						e += "s"
					}
				} else if cellWidth >= 5 {
					e = "󱑑"
				} else {
					e = ""
				}
				rows[row][col] += s.Render(" " + strconv.Itoa(nevents) + " " + e)

				// TODO: Skip events that are finished
				if cellWidth <= 8 {
					continue
				}

				events := calendar.GetEventsByDate(
					*m.focusYear,
					*m.focusMonth,
					day,
					ptrCalendarsToCalendars(m.calendars),
				)

				if row+1 == hoverRow && col == hoverCol {
					s = s.UnsetForeground().Foreground(purple.Colors.LightGray)
				}

				for _, e := range events {
					rows[row][col] += "\n"
					if len("- "+e.Name) > cellWidth {
						rows[row][col] += s.Render(" " + e.Name[:cellWidth-5] + "...")
					} else {
						rows[row][col] += s.Render(" " + e.Name)
					}
				}
			}
		}
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(style.BorderStyle).
		BorderRow(true).
		BorderHeader(true).
		Headers(getHeaders(m.width)...).
		Rows(rows...).
		Width(m.width).Height(m.height - 1).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return style.TitleStyle
			}
			var s lipgloss.Style
			if row == hoverRow && col == hoverCol {
				s = style.CellStyleHover
			} else {
				s = style.CellStyle
			}

			if row == todayRow && col == todayCol {
				s = s.UnsetForeground().Foreground(purple.Colors.Accent).SetString("󰃭")
			}

			return s.MaxHeight(cellHeight).Height(cellHeight).Width(cellWidth)
		})

	return t.Render() + "\n"
}

func (m Model) drawTitle() string {
	return style.TitleStyle.Width(m.width).
		Render(fmt.Sprintf("󰸗 %s %d", time.Month(*m.focusMonth).String(), *m.focusYear)) + "\n"
}

func (m Model) drawNotice() string {
	return style.Notice.Width(m.width).
		Render(fmt.Sprintf(" %d events this month", calendar.GetNumberOfEventsInMonth(*m.focusYear, *m.focusMonth, ptrCalendarsToCalendars(m.calendars))))
}

func (m Model) View() string {
	var str string
	str += m.drawTitle()
	str += m.drawCalendar()
	str += m.drawNotice()
	return lipgloss.NewStyle().Height(m.height).MaxHeight(m.height).Width(m.width).Render(str)
}
