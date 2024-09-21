package newevent

import (
	"fmt"
	"regexp"

	"github.com/anotherhadi/calendar/style"
	calendar "github.com/anotherhadi/markdown-calendar"
	"github.com/charmbracelet/huh"
)

type Model struct {
	calendars *calendar.Calendar
	form      *huh.Form

	background   string
	previousView string

	width, height int
}

const dateFormat = "DD/MM/YYYY"

func NewModel(
	calendars *calendar.Calendar,
	background string,
	previousView string,
) Model {
	form := huh.NewForm(
		huh.NewGroup(
			// TODO: make this one optional if there are only 1 calendar
			// huh.NewSelect[string]().
			// 	Key("calendar").
			// 	Options(huh.NewOptions(calendar.GetCalendarsNames(utils.PtrCalendarsToCalendars(calendar.GetPurpleCalendars()))...)...).
			// 	Title("Choose a calendar"),

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
	form.WithTheme(style.GetFormTheme())

	m := Model{
		form:      form,
		calendars: calendars,

		background:   background,
		previousView: previousView,

		width:  0,
		height: 0,
	}

	return m
}
