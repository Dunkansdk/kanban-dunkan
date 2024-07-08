package footer

import tea "github.com/charmbracelet/bubbletea"

type FooterMsg struct {
	mode string
}

func UpdateFooterMode(mode string) tea.Msg {
	return FooterMsg{mode: mode}
}
