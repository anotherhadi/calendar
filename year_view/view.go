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
			Render("◖")
		str += lipgloss.NewStyle().
			Background(purple.Colors.Accent).
			Foreground(purple.GetFgColor(purple.Colors.Accent)).
			Align(lipgloss.Center).
			Width(20 - 2).
			Render(time.Month(month).String())
		str += lipgloss.NewStyle().
			Foreground(purple.Colors.Accent).
			Render("◗")
		str += "\n"
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

	if m.height > 40 {
		str += lipgloss.NewStyle().
			Foreground(purple.Colors.Muted).
			Render("Su Mo Tu We Th Fr Sa") +
			"\n"
	}
	if m.height > 50 {
		str += "\n"
	}

	for i := 1; i <= daysInMonth; i++ {
		if i == 1 {
			for j := 0; j < startDay; j++ {
				str += "   "
			}
		}

		var s lipgloss.Style

		if i == m.currentDay && month == m.currentMonth && year == m.currentYear {
			s = lipgloss.NewStyle().
				Background(purple.Colors.Accent).
				Foreground(purple.GetFgColor(purple.Colors.Accent))
		} else if len(m.calendar.GetEventsByDate(year, month, i)) > 0 {
			s = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
		} else {
			s = lipgloss.NewStyle()
		}

		str += s.Render(fmt.Sprintf("%2d ", i))
		if (i+startDay)%7 == 0 {
			str += "\n"
		}
	}

	s := lipgloss.NewStyle()
	if m.height > 34 {
		s = s.MarginTop(1)
	} else {
		s = s.MarginTop(0).MarginBottom(0)
	}
	if m.width > 100 {
		s = s.MarginLeft(4).MarginRight(4)
	} else if m.width > 70 {
		s = s.MarginLeft(2).MarginRight(2)
	} else {
		s = s.MarginLeft(1).MarginRight(0)
	}

	return s.Render(str)
}

func (m Model) drawYear() string {
	var str string
	var maxWidth int = 120
	var currentLine string
	width := m.width
	if width > maxWidth {
		width = maxWidth
	}

	for i := 1; i <= 12; i++ {
		y := m.drawMinimalCalendar(i, *m.focusYear)
		if len(
			utils.RemoveAnsiStyle(
				strings.Split(lipgloss.JoinHorizontal(lipgloss.Top, currentLine, y), "\n")[0],
			),
		) > width {
			str = lipgloss.JoinVertical(lipgloss.Top, str, currentLine)
			currentLine = y
		} else {
			currentLine = lipgloss.JoinHorizontal(lipgloss.Top, currentLine, y)
		}
	}
	str = lipgloss.JoinVertical(lipgloss.Top, str, currentLine)

	return lipgloss.Place(m.width, m.height-2, lipgloss.Center, lipgloss.Center, str) + "\n"
}

func (m Model) drawTitle() string {
	return style.TitleStyle.Width(m.width).
		Render(fmt.Sprintf(utils.YearIcon+" Year %d", *m.focusYear)) +
		"\n"
}

func (m Model) drawNotice() string {
	return style.Notice.Width(m.width).
		Render(fmt.Sprintf(utils.NoticeIcon+" %d events this year", len(m.calendar.GetEventsByYear(*m.focusYear))))
}

func (m Model) View() string {
	var str string
	str += m.drawTitle()
	str += m.drawYear()
	str += m.drawNotice()

	return lipgloss.NewStyle().Height(m.height).MaxHeight(m.height).Width(m.width).Render(str)
}
