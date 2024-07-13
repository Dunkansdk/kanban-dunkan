package navigation

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ModelRestoreMsg struct{}
type ModelPopMsg struct{}
type ModelPushMsg struct {
	Item NavigationItem
}
type ModelReplaceMsg struct {
	Item NavigationItem
}

func Pop() tea.Cmd {
	return func() tea.Msg {
		return ModelPopMsg{}
	}
}

func Push(item NavigationItem) tea.Cmd {
	return func() tea.Msg {
		return ModelPushMsg{item}
	}
}

func Replace(item NavigationItem) tea.Cmd {
	return func() tea.Msg {
		return ModelReplaceMsg{item}
	}
}
