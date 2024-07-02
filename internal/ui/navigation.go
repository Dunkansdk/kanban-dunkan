package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type NavigationStack struct {
	// zones *zone.Manager
	stack []tea.Model
	size  *tea.WindowSizeMsg
}

func NewNavigation(root tea.Model) NavigationStack {
	navigation := NavigationStack{
		stack: []tea.Model{root},
	}
	return navigation
}

func (navigation NavigationStack) Init() tea.Cmd {
	return navigation.initRoot()
}

func (navigation NavigationStack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ModelPopMsg:
		return navigation, navigation.Pop()
	case ModelPushMsg:
		return navigation, navigation.Push(msg.Page)
	case ModelReplaceMsg:
		return navigation, navigation.Replace(msg.Page)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return navigation, navigation.Pop()
		}
	case tea.WindowSizeMsg:
		navigation.size = &msg
	}
	var cmd tea.Cmd
	navigation.stack[len(navigation.stack)-1], cmd = navigation.Top().Update(msg)
	return navigation, cmd
}

func (navigation NavigationStack) View() string {
	return navigation.Top().View()
}
