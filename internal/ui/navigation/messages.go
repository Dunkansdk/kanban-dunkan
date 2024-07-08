package navigation

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ModelRestoreMsg struct{}
type ModelPopMsg struct{}
type ModelPushMsg struct {
	Page tea.Model
}
type ModelReplaceMsg struct {
	Page tea.Model
}

func Pop() tea.Cmd {
	return func() tea.Msg {
		return ModelPopMsg{}
	}
}

func Push(page tea.Model) tea.Cmd {
	return func() tea.Msg {
		return ModelPushMsg{page}
	}
}

func Replace(page tea.Model) tea.Cmd {
	return func() tea.Msg {
		return ModelReplaceMsg{page}
	}
}
