package month

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/anotherhadi/calendar/style"
	"github.com/anotherhadi/calendar/utils"
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

	calendars := calendar.GetPurpleCalendars()
	eventColors := make(map[string]string, len(calendars))
	for i := range calendars {
		eventColors[calendars[i].Name] = calendars[i].EventColor
	}

	if cellHeight > 1 {
		for row := range rows {
			for col := range rows[row] {
				day, _ := strconv.Atoi(rows[row][col])
				if row == 0 && day > 7 {
					continue
				} else if row == len(rows)-1 && day < 7 {
					continue
				}

				nevents := len(m.calendar.GetEventsByDate(*m.focusYear, *m.focusMonth, day))
				rows[row][col] += "\n"

				if nevents < 1 {
					continue
				}
				if cellWidth < 4 {
					continue
				}

				var s lipgloss.Style

				s = style.EventStyle
				if row+1 == hoverRow && col == hoverCol {
					s = s.UnsetForeground().Foreground(purple.Colors.Accent)
					s = s.UnsetBackground().Background(purple.Colors.Muted)
				}

				// TODO: Skip events that are finished if not enough space to show all
				events := m.calendar.GetEventsByDate(*m.focusYear, *m.focusMonth, day)

				for i, e := range events {
					if i >= cellHeight-2 {
						remaining := nevents - i
						if remaining != 1 {
							rows[row][col] += s.Foreground(purple.Colors.LightGray).
								Width(cellWidth).MaxHeight(1).Align(lipgloss.Center).
								Render(utils.TruncateString("+"+strconv.Itoa(remaining), cellWidth-2))
							rows[row][col] += "\n"
							break
						}
					}

					rows[row][col] += s.Foreground(lipgloss.Color(eventColors[e.CalendarName])).
						Render("◖")
					rows[row][col] += s.Background(lipgloss.Color(eventColors[e.CalendarName])).
						Foreground(purple.GetFgColor(lipgloss.Color(eventColors[e.CalendarName]))).
						Width(cellWidth - 2).MaxHeight(1).
						Render(utils.TruncateString(drawEvent(e), cellWidth-2))
					rows[row][col] += s.Foreground(lipgloss.Color(eventColors[e.CalendarName])).
						Render("◗")
					rows[row][col] += "\n"
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
		Width(m.width).Height(m.height - 2).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return style.TitleStyle
			}
			var s lipgloss.Style
			day, _ := strconv.Atoi(strings.Split(rows[row-1][col], "\n")[0])
			if row == hoverRow && col == hoverCol {
				s = style.CellStyleHover
			} else if row == 1 && day > 7 {
				s = style.OutsideCellStyle
			} else if row == len(rows) && day < 7 {
				s = style.OutsideCellStyle
			} else {
				s = style.CellStyle
			}

			if row == todayRow && col == todayCol {
				s = s.UnsetForeground().Foreground(purple.Colors.Accent).SetString(utils.TodayIcon)
			}

			return s.MaxHeight(cellHeight).Height(cellHeight).Width(cellWidth)
		})

	return t.Render() + "\n"
}

func drawEvent(event calendar.Event) string {
	prefix := ""
	if event.AllDay {
		prefix = utils.AllDayIcon + " | "
	} else {
		prefix = fmt.Sprintf("%02d:%02d | ", event.StartDate.Hour, event.StartDate.Minute)
	}

	return prefix + event.Name
}

func (m Model) drawTitle() string {
	return style.TitleStyle.Width(m.width).
		Render(fmt.Sprintf(utils.MonthIcon+" %s %d", time.Month(*m.focusMonth).String(), *m.focusYear)) +
		"\n"
}

func (m Model) drawNotice() string {
	return style.Notice.Width(m.width).
		Render(fmt.Sprintf(utils.NoticeIcon+" %d events this month", len(m.calendar.GetEventsByMonth(*m.focusYear, *m.focusMonth))))
}

func (m Model) View() string {
	var str string
	str += m.drawTitle()
	str += m.drawCalendar()
	str += m.drawNotice()
	return lipgloss.NewStyle().Height(m.height).MaxHeight(m.height).Width(m.width).Render(str)
}
