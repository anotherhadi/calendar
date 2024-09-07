package monthview

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	TitleStyle       = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Align(lipgloss.Center).Bold(true)
	HeaderStyle      = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Align(lipgloss.Center).Bold(true)
	BorderStyle      = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	CellStyle        = lipgloss.NewStyle().Padding(0, 1)
	TodayCellStyle   = CellStyle.Background(purple.Colors.Muted)
	OutsideCellStyle = CellStyle.Foreground(purple.Colors.Muted)
	EventStyle       = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	EventStyleHover  = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	Notice           = lipgloss.NewStyle().Foreground(purple.Colors.Muted).Align(lipgloss.Center)
)

func (m Model) drawCalendar() string {
	rows := [][]string{}

	var hoverRow, hoverCol int
	var todayRow, todayCol int

	for i := 1; i <= calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear); i++ {
		if i == 1 {
			tmp := []string{}
			for j := 0; j < calendar.DayOfWeek(i, *m.FocusMonth, *m.FocusYear); j++ {
				// Place the previous month's days in the calendar
				tmp = append(tmp, OutsideCellStyle.Render(strconv.Itoa(calendar.DaysInMonth(*m.FocusMonth-1, *m.FocusYear)-calendar.DayOfWeek(i, *m.FocusMonth, *m.FocusYear)+j+1)))
			}
			rows = append(rows, tmp)
		}

		if i == *m.FocusDay {
			hoverRow = len(rows)
			hoverCol = len(rows[len(rows)-1])
		}

		if i == m.CurrentDay && *m.FocusMonth == m.CurrentMonth && *m.FocusYear == m.CurrentYear {
			todayRow = len(rows)
			todayCol = len(rows[len(rows)-1])
			rows[len(rows)-1] = append(rows[len(rows)-1], today(strconv.Itoa(i), i == *m.FocusDay))
		} else {
			rows[len(rows)-1] = append(rows[len(rows)-1], strconv.Itoa(i))
		}

		if calendar.DayOfWeek(i, *m.FocusMonth, *m.FocusYear) == 6 {
			rows = append(rows, []string{})
		}
	}

	// Delete the last row if it's empty
	if len(rows[len(rows)-1]) == 0 {
		rows = rows[:len(rows)-1]
	}
	// Place the next month's days in the calendar
	for i := 1; len(rows[len(rows)-1]) < 7; i++ {
		rows[len(rows)-1] = append(rows[len(rows)-1], OutsideCellStyle.Render(strconv.Itoa(i)))
	}

	nrow := len(rows)
	heightAvailablePerCell := int((m.Height - 1 - 3 - nrow) / nrow) // 1 for title, 3 for header, nrow for days border
	if heightAvailablePerCell > 1 {
		for row := range rows {
			for col := range rows[row] {
				if rows[row][col] == "" {
					continue
				}
				day, err := strconv.Atoi(rows[row][col])
				if todayRow == row+1 && todayCol == col {
					day = m.CurrentDay
					err = nil
				}
				if err != nil {
					continue
				}
				nevents := m.getNumberOfEvents(*m.FocusYear, *m.FocusMonth, day)
				s := EventStyle
				if row+1 == hoverRow && col == hoverCol {
					s = EventStyleHover
				}
				rows[row][col] += "\n"
				if nevents == 1 {
					rows[row][col] += s.Render(" " + strconv.Itoa(nevents) + " event")
				} else if nevents > 1 {
					rows[row][col] += s.Render(" " + strconv.Itoa(nevents) + " events")
				}
				if heightAvailablePerCell > 2 {
					rows[row][col] += strings.Repeat("\n", heightAvailablePerCell-1-1) // -1 for the day number and -1 for the events notice
				}
			}
		}
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(BorderStyle).
		BorderRow(true).
		BorderHeader(true).
		Headers(getHeaders(m.Width)...).
		Rows(rows...).
		Width(m.Width).Height(m.Height - 1).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return HeaderStyle
			}
			if row == hoverRow && col == hoverCol {
				return TodayCellStyle
			}
			return CellStyle
		})

	var str string
	str += TitleStyle.Width(m.Width).Render(fmt.Sprintf("󰸗 %s %d", time.Month(*m.FocusMonth).String(), *m.FocusYear)) + "\n"
	str += t.Render()

	tableHeight := strings.Count(str, "\n") + 1
	if tableHeight < m.Height {
		str += strings.Repeat("\n", m.Height-tableHeight)
		str += Notice.Width(m.Width).Render(fmt.Sprintf(" %d events this month", m.getNumberOfEventsInMonth(*m.FocusYear, *m.FocusMonth)))
	}

	return str
}
