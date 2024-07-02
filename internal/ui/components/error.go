package components

import "github.com/charmbracelet/lipgloss"

func Error(width int, height int, msg string) string {
	err := lipgloss.NewStyle().Width(width-30).Border(lipgloss.ThickBorder(), true).Padding(0, 2, 0, 1).BorderForeground(lipgloss.Color("7")).Foreground(lipgloss.Color("15")).Render(msg)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, err, lipgloss.WithWhitespaceChars("!"), lipgloss.WithWhitespaceForeground(lipgloss.Color("7")))
}
