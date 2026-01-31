package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)


func (m model) View() string {
	var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15")).Padding(1, 1).MarginLeft(28).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("39"))

	var helpStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("202"))


	welcome := style.Render("Welcome to App 🫶")

	help := helpStyle.Render("Ctrl+N: new file   *   Ctrl+L: list   *   Esc: back/save   *   Ctrl+S: save   *   Ctrl+Q: quit")

	view := ""

	if m.visibleInputText {
		view = m.newFileInput.View()
	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome,view,help)
}