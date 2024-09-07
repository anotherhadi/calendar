package main

func (m model) View() string {
	var str string
	if m.CurrentView == "month" {
		str = m.MonthModel.View()
	}
	return str
}
