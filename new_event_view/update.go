package newevent

import (
	"github.com/anotherhadi/calendar/utils"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/huh"
)

func (m Model) Update(message tea.Msg) (Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.form.WithWidth(msg.Width)
		m.form.WithHeight(msg.Height)
		return m, nil
	}

	// form, cmd := m.form.Update(message)
	form, _ := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		// 	m.Event.Name = m.form.GetString("name")
		// 	m.Event.Description = m.form.GetString("description")
		// 	m.Event.StartDate.Day, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[0])
		// 	m.Event.StartDate.Month, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[1])
		// 	m.Event.StartDate.Year, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[2])
		// 	for i := range m.calendars {
		// 		if m.calendars[i].Name == *m.CalendarName {
		// 			m.calendars[i].AddEvent(*m.Event)
		// 			_ = m.calendars[i].Write()
		// 			*m.isViewed = false
		// 			return m, nil
		// 		}
		// 	}
		return m, utils.ChangeFocusViewCmd(m.previousView)
	}

	// return m, cmd
	return m, nil
}
