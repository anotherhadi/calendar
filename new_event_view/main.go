package newevent

import (
	"github.com/anotherhadi/calendar/utils"
	tea "github.com/charmbracelet/bubbletea/v2"
)

func (m Model) Init() (Model, tea.Cmd) {
	return m, utils.WrapOldBubbleteaCmd(m.form.Init())
	// return m, nil
}
