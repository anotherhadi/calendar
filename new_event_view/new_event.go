package neweventview

import (
	"fmt"
	"regexp"

	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	form      *huh.Form
	Event     *calendar.Event
	calendars []*calendar.Calendar
}

var (
	dateFormat = "dd/mm/yyyy"
	TitleStyle = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Bold(true).Underline(true)
)

func getFormTheme() *huh.Theme {
	theme := huh.ThemeBase()
	theme.Focused.Title = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Blurred.Title = lipgloss.NewStyle().Foreground(purple.Colors.LightGray)
	theme.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true).BorderForeground(purple.Colors.Accent)
	theme.Blurred.Description = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	theme.Focused.Description = lipgloss.NewStyle().Foreground(purple.Colors.Muted)
	theme.Focused.TextInput.Prompt = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Blurred.TextInput.Prompt = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	theme.Focused.SelectSelector = lipgloss.NewStyle().Foreground(purple.Colors.Muted).SetString("> ")
	theme.Focused.SelectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Accent)
	theme.Focused.UnselectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	theme.Blurred.SelectSelector = lipgloss.NewStyle().Foreground(purple.Colors.Muted).SetString("> ")
	theme.Blurred.SelectedOption = lipgloss.NewStyle().Foreground(purple.Colors.LightGray)
	theme.Blurred.UnselectedOption = lipgloss.NewStyle().Foreground(purple.Colors.Muted)

	return theme
}

func NewModel(calendars []*calendar.Calendar) Model {
	calendarsName := make([]string, len(calendars))
	for i, c := range calendars {
		calendarsName[i] = c.Name
	}

	event := calendar.Event{}

	form := huh.NewForm(
		huh.NewGroup(
			// TODO: make this one optional if there are only 1 calendar
			huh.NewSelect[string]().
				Key("calendar").
				Options(huh.NewOptions(calendarsName...)...).
				Title("Choose a calendar"),

			huh.NewInput().
				Title("Name").
				Key("name"),

			huh.NewInput().
				Validate(func(s string) error {
					if !regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`).MatchString(s) {
						return fmt.Errorf("Invalid date format: " + dateFormat)
					}
					return nil
				}).
				Title("Date").
				Description(dateFormat).
				Key("date"),

			// TODO: All-day event, dynamic fields for end date and time

			huh.NewText().
				Title("Description").
				Key("description"),
		),
	)
	form.WithWidth(0)
	form.WithHeight(0)
	form.WithTheme(getFormTheme())
	return Model{form: form, Event: &event, calendars: calendars}
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func FormSent() tea.Msg {
	return FormSentMsg{}
}

type FormSentMsg struct{}

func (m Model) Update(message tea.Msg) (Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.WindowSizeMsg:
		m.form.WithWidth(msg.Width - 10)
		m.form.WithHeight(msg.Height - 5)
		return m, nil
	}

	form, cmd := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	// if m.form.State == huh.StateCompleted {
	// 	return m, FormSent
	// }

	return m, cmd
}

func (m Model) View() string {
	var str string
	str += TitleStyle.Render("# New Event") + "\n\n"
	str += m.form.View()
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(purple.Colors.Accent).Padding(1, 5).Render(str)
}
