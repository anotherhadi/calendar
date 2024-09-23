package month

import (
	"github.com/anotherhadi/calendar/utils"
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/charmbracelet/bubbles/v2/key"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func (m Model) Update(message tea.Msg) (Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// case tea.MouseMsg:
	// 	mouse := msg.Mouse()
	// 	if len(*m.rawCalendar) == 0 {
	// 		return m, nil
	// 	}
	// 	switch msg := msg.(type) {
	// 	case tea.MouseClickMsg:
	// 		switch msg.Button {
	// 		case tea.MouseLeft:
	// 			x := mouse.X - 1
	// 			y := mouse.Y - 4
	// 			calWidth := m.width - 2                                // ignore left & right borders
	// 			calHeight := strings.Count(m.drawCalendar(), "\n") - 4 // ignore top & bottom borders & headers
	// 			if x < 0 || x > calWidth || y < 0 || y > calHeight {
	// 				return m, nil
	// 			}
	//
	// 			// Calculate the day clicked
	//
	// 			cellWidth := int((calWidth - 7) / 7)
	// 			cellHeight := int(
	// 				(calHeight - len(*m.rawCalendar)) / len(*m.rawCalendar),
	// 			)
	//
	// 			colClicked := int(x / (cellWidth))
	// 			rowClicked := int(y / (cellHeight))
	//
	// 			// rowClicked := int(y / (calHeight / (len(*m.rawCalendar) - 1)))
	// 			// colClicked := int(x / ((calWidth / 7) + 1))
	// 			// dayClicked, _ := strconv.Atoi((*m.rawCalendar)[rowClicked][colClicked])
	// 			// *m.focusDay = dayClicked
	// 			*m.msg = fmt.Sprintf("row:%d col:%d", rowClicked, colClicked)
	//
	// 			return m, nil
	// 		}
	// 	}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.PreviousWeek):
			calendar.IncrementDay(m.focusDay, m.focusMonth, m.focusYear, -7)
		case key.Matches(msg, m.keys.NextWeek):
			calendar.IncrementDay(m.focusDay, m.focusMonth, m.focusYear, 7)
		case key.Matches(msg, m.keys.PreviousDay):
			calendar.IncrementDay(m.focusDay, m.focusMonth, m.focusYear, -1)
		case key.Matches(msg, m.keys.NextDay):
			calendar.IncrementDay(m.focusDay, m.focusMonth, m.focusYear, 1)
		case key.Matches(msg, m.keys.PreviousYear):
			calendar.IncrementYear(m.focusDay, m.focusMonth, m.focusYear, -1)
		case key.Matches(msg, m.keys.NextYear):
			calendar.IncrementYear(m.focusDay, m.focusMonth, m.focusYear, 1)
		case key.Matches(msg, m.keys.PreviousMonth):
			calendar.IncrementMonth(m.focusDay, m.focusMonth, m.focusYear, -1)
		case key.Matches(msg, m.keys.NextMonth):
			calendar.IncrementMonth(m.focusDay, m.focusMonth, m.focusYear, 1)
		case key.Matches(msg, m.keys.NewEvent):
			return m, utils.ChangeFocusViewCmd("new_event")
		case key.Matches(msg, m.keys.DayView):
			return m, utils.ChangeFocusViewCmd("day")
		case key.Matches(msg, m.keys.WeekView):
			return m, utils.ChangeFocusViewCmd("week")
		case key.Matches(msg, m.keys.YearView):
			return m, utils.ChangeFocusViewCmd("year")
		case key.Matches(msg, m.keys.Today):
			*m.focusDay, *m.focusMonth, *m.focusYear = m.currentDay, m.currentMonth, m.currentYear
		case key.Matches(msg, m.keys.Help):
			m.Help.ShowAll = !m.Help.ShowAll
		}
	}

	return m, nil
}
