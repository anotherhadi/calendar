package year

import (
	"fmt"
	"strings"
	"time"

	"github.com/anotherhadi/calendar/style"
	"github.com/anotherhadi/calendar/utils"
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) drawMinimalCalendar(month, year int) string {
	var str string
	if month == *m.focusMonth && year == *m.focusYear {
		str += lipgloss.NewStyle().
			Foreground(purple.Colors.Accent).
			Align(lipgloss.Center).
			Width(20).
			Render(time.Month(month).String()) +
			"\n"
	} else {
		str += lipgloss.NewStyle().
			Foreground(purple.Colors.LightGray).
			Align(lipgloss.Center).
			Width(20).
			Render(time.Month(month).String()) +
			"\n"
	}

	daysInMonth := calendar.DaysInMonth(month, year)
	startDay := calendar.DayOfWeek(1, month, year)

	// Draw the days of the week
	str += lipgloss.NewStyle().Foreground(purple.Colors.Muted).Render("Su Mo Tu We Th Fr Sa") + "\n"

	// Draw the rest of the weeks
	for i := 1; i <= daysInMonth; i++ {
		if i == 1 {
			for j := 0; j < startDay; j++ {
				str += "   "
			}
		}

		var s lipgloss.Style

		if i == m.currentDay && month == m.currentMonth && year == m.currentYear {
			s = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
		} else {
			s = lipgloss.NewStyle()
		}

		str += s.Render(fmt.Sprintf("%2d ", i))
		if (i+startDay)%7 == 0 {
			str += "\n"
		}
	}

	return lipgloss.NewStyle().Margin(1, 2, 0, 2).Render(str)
}

func (m Model) drawYear() string {
	var str string
	var currentLine string

	for i := 1; i <= 12; i++ {
		y := m.drawMinimalCalendar(i, *m.focusYear)
		if len(
			utils.RemoveAnsiStyle(
				strings.Split(lipgloss.JoinHorizontal(lipgloss.Top, currentLine, y), "\n")[0],
			),
		) > m.width {
			str = lipgloss.JoinVertical(lipgloss.Top, str, currentLine)
			currentLine = y
		} else {
			currentLine = lipgloss.JoinHorizontal(lipgloss.Top, currentLine, y)
		}
	}
	str = lipgloss.JoinVertical(lipgloss.Top, str, currentLine)

	return lipgloss.Place(m.width, m.height-2, lipgloss.Center, lipgloss.Center, str)
}

func (m Model) drawTitle() string {
	return style.TitleStyle.Width(m.width).
		Render(fmt.Sprintf(utils.YearIcon+" Year %d", *m.focusYear)) +
		"\n"
}

func (m Model) View() string {
	var str string
	str += m.drawTitle()
	str += m.drawYear()

	return lipgloss.NewStyle().Height(m.height).MaxHeight(m.height).Width(m.width).Render(str)
}
