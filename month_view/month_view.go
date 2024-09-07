package monthview

import (
	calendar "github.com/anotherhadi/markdown-calendar"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	CurrentDay, CurrentMonth, CurrentYear int
	FocusDay, FocusMonth, FocusYear       *int

	Calendars []*calendar.Calendar

	Width, Height int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) IncrementYear(inc int) {
	*m.FocusYear += inc
	if *m.FocusYear < 1 {
		*m.FocusYear = 1
	}
	if *m.FocusYear > 9999 {
		*m.FocusYear = 9999
	}
	if *m.FocusDay > calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear) {
		*m.FocusDay = calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear)
	}
}

func (m *Model) IncrementMonth(inc int) {
	*m.FocusMonth += inc
	if *m.FocusMonth < 1 {
		*m.FocusMonth = 12
		m.IncrementYear(-1)
	}
	if *m.FocusMonth > 12 {
		*m.FocusMonth = 1
		m.IncrementYear(1)
	}
	if *m.FocusDay > calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear) {
		*m.FocusDay = calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear)
	}
}

func (m *Model) IncrementDay(inc int) {
	focusDay := *m.FocusDay
	focusDay += inc
	if focusDay < 1 {
		m.IncrementMonth(-1)
		*m.FocusDay = calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear) + inc + *m.FocusDay
	} else if focusDay > calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear) {
		*m.FocusDay = *m.FocusDay + inc - calendar.DaysInMonth(*m.FocusMonth, *m.FocusYear)
		m.IncrementMonth(1)
	} else {
		*m.FocusDay = focusDay
	}
}

func (m Model) Update(message tea.Msg) (Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k": // Previous week
			m.IncrementDay(-7)
		case "down", "j": // Next week
			m.IncrementDay(7)
		case "left", "h": // Previous day
			m.IncrementDay(-1)
		case "right", "l": // Next day
			m.IncrementDay(1)
		case "K": // Previous year
			m.IncrementYear(-1)
		case "J": // Next year
			m.IncrementYear(1)
		case "H": // Previous month
			m.IncrementMonth(-1)
		case "L": // Next month
			m.IncrementMonth(1)
		}
	}

	return m, nil
}

func (m Model) View() string {
	var str string
	str += m.drawCalendar()
	return str
}
