package navigation

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (navigation NavigationStack) initRoot() tea.Cmd {
	cmds := []tea.Cmd{navigation.Top().Init()}
	if navigation.size != nil {
		top, cmd := navigation.Top().Update(*navigation.size)
		navigation.stack[len(navigation.stack)-1] = top
		cmds = append(cmds, cmd)
	}
	return tea.Batch(cmds...)
}

func (navigation NavigationStack) Top() tea.Model {
	if len(navigation.stack) == 0 {
		return nil
	}
	return navigation.stack[len(navigation.stack)-1]
}

func (navigation *NavigationStack) Push(model tea.Model) tea.Cmd {
	navigation.stack = append(navigation.stack, model)
	return navigation.initRoot()
}

func (navigation *NavigationStack) Replace(model tea.Model) tea.Cmd {
	navigation.stack[len(navigation.stack)-1] = model
	return navigation.initRoot()
}

func (navigation *NavigationStack) Pop() tea.Cmd {
	if len(navigation.stack) == 1 {
		return tea.Quit
	}
	navigation.stack = navigation.stack[:len(navigation.stack)-1]
	if navigation.size == nil {
		return nil
	}
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	navigation.stack[len(navigation.stack)-1], cmd = navigation.Top().Update(*navigation.size)
	cmds = append(cmds, cmd)
	navigation.stack[len(navigation.stack)-1], cmd = navigation.Top().Update(ModelRestoreMsg{})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
