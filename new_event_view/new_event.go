package neweventview

import (
	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Bold(true).Underline(true)

	focusedStyle = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	blurredStyle = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Padding(0, 1).Border(lipgloss.RoundedBorder()).BorderForeground(purple.Colors.Accent).Render("Create")
	blurredButton = blurredStyle.Padding(0, 1).Border(lipgloss.RoundedBorder()).BorderForeground(purple.Colors.Muted).Render("Create")
)

type Model struct {
	focusIndex int
	Inputs     []textinput.Model
	Labels     []string
	Title      string

	Width, Height int
}

func InitialModel(title string) Model {
	m := Model{
		Inputs: make([]textinput.Model, 3),
		Labels: make([]string, 3),
		Title:  title,
	}

	var t textinput.Model
	for i := range m.Inputs {
		t = textinput.New()
		t.Prompt = " "
		t.PromptStyle = blurredStyle
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			m.Labels[i] = "Event name"
			t.Placeholder = "My cool event"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			m.Labels[i] = "Date"
			t.Placeholder = "yyyy-mm-dd"
			t.CharLimit = 10
		case 2:
			m.Labels[i] = "Description"
			t.Placeholder = "The event description"
			t.CharLimit = 100
		}

		m.Inputs[i] = t
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			// if s == "enter" && m.focusIndex == len(m.inputs) {
			// 	return m, tea.Quit
			// }

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.Inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.Inputs)
			}

			cmds := make([]tea.Cmd, len(m.Inputs))
			for i := 0; i <= len(m.Inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.Inputs[i].Focus()
					m.Inputs[i].PromptStyle = focusedStyle
					m.Inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.Inputs[i].Blur()
				m.Inputs[i].PromptStyle = blurredStyle
				m.Inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.Inputs))

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var str string

	str += blurredStyle.Render("# ") + TitleStyle.Render(m.Title) + "\n\n"

	for i := range m.Inputs {
		if i == m.focusIndex {
			str += focusedStyle.Render(m.Labels[i] + ":")
		} else {
			str += blurredStyle.Render(m.Labels[i] + ":")
		}
		str += "\n"
		str += m.Inputs[i].View()
		if i < len(m.Inputs)-1 {
			str += "\n\n"
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.Inputs) {
		button = &focusedButton
	}
	str += "\n\n" + *button
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(purple.Colors.Accent).Width(m.Width).Height(m.Height).Padding(1, 5).Render(str)
}
