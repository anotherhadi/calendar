package neweventview

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/anotherhadi/purple-apps"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Bold(true).Underline(true)
	AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))

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
	dateFormat := "dd/mm/yyyy"
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
			m.Labels[i] = "Date (" + dateFormat + ")"
			t.Placeholder = dateFormat
			t.CharLimit = 10
			t.Validate = func(value string) error {

				goodMatch := "08/09/2024"
				var valueCompleted string = value
				if len(value) != len(goodMatch) {
					valueCompleted += goodMatch[len(value):]
				}
				if !regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`).MatchString(valueCompleted) {
					return errors.New("Invalid date format: " + dateFormat + valueCompleted)
				}
				if value == "" {
					return nil
				}
				splitted := strings.Split(value, "/")

				if len(splitted) >= 1 {
					day, _ := strconv.Atoi(splitted[0])
					if day > 31 || day < 1 && splitted[0] != "0" {
						return errors.New("Invalid day: " + dateFormat)
					}
				}
				if len(splitted) >= 2 {
					month, _ := strconv.Atoi(splitted[1])
					if month > 12 || month < 1 && splitted[1] != "" {
						return errors.New("Invalid month: " + dateFormat)
					}
				}
				if len(splitted) >= 3 {
					year, _ := strconv.Atoi(splitted[2])
					if year < 0 {
						return errors.New("Invalid year: " + dateFormat)
					}
				}

				return nil
			}
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
		if m.Inputs[i].Err != nil {
			str += "\n" + AlertStyle.Render(m.Inputs[i].Err.Error())
		}
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
