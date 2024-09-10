package neweventview

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/anotherhadi/purple-apps"
	tea "github.com/charmbracelet/bubbletea"
	huh "github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	form          *huh.Form
	Event         *calendar.Event
	CalendarName  *string
	calendars     []*calendar.Calendar
	isViewed      *bool
	Width, Height int
}

const dateFormat = "DD/MM/YYYY"

var (
	TitleStyle = lipgloss.NewStyle().Foreground(purple.Colors.Accent).Bold(true).Underline(true)
)

func NewModel(calendars []*calendar.Calendar, isViewed *bool, day, month, year int) Model {
	calendarsName := make([]string, len(calendars))
	for i, c := range calendars {
		calendarsName[i] = c.Name
	}

	event := &calendar.Event{}
	calendarName := ""
	date := fmt.Sprintf("%02d/%02d/%04d", day, month, year)

	form := huh.NewForm(
		huh.NewGroup(
			// TODO: make this one optional if there are only 1 calendar
			huh.NewSelect[string]().
				Key("calendar").
				Options(huh.NewOptions(calendarsName...)...).
				Title("Choose a calendar").
				Value(&calendarName),

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
				Value(&date).
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
	return Model{form: form, calendars: calendars, Event: event, CalendarName: &calendarName, isViewed: isViewed}
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
		m.Width = msg.Width
		m.Height = msg.Height
		m.form.WithWidth(msg.Width - 20) // Padding * 2 + borders + margins
		m.form.WithHeight(msg.Height - 11)
		return m, nil
	}

	form, cmd := m.form.Update(message)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	if m.form.State == huh.StateCompleted {
		m.Event.Name = m.form.GetString("name")
		m.Event.Description = m.form.GetString("description")
		m.Event.StartDate.Day, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[0])
		m.Event.StartDate.Month, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[1])
		m.Event.StartDate.Year, _ = strconv.Atoi(strings.Split(m.form.GetString("date"), "/")[2])
		for i := range m.calendars {
			if m.calendars[i].Name == *m.CalendarName {
				m.calendars[i].AddEvent(*m.Event)
				_ = m.calendars[i].Write()
				*m.isViewed = false
				return m, nil
			}
		}
	}

	return m, cmd
}

func (m Model) View() string {
	var str string

	title := TitleStyle.Render("# New Event") + "\n\n"
	form := m.form.View()
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).BorderForeground(purple.Colors.Accent).
		Padding(1, 5).
		Render(title + form)

	str += lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, box,
		lipgloss.WithWhitespaceChars("/"), lipgloss.WithWhitespaceForeground(purple.Colors.Muted))

	return str
}
