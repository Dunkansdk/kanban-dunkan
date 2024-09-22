package navigation

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (navigation NavigationStack) initRoot() tea.Cmd {
	cmds := []tea.Cmd{navigation.Top().Init()}
	if navigation.size != nil {
		model, cmd := navigation.Top().Update(*navigation.size)
		navigation.stack[len(navigation.stack)-1].Model = model
		cmds = append(cmds, cmd)
	}
	navigation.footer.UpdateContent(navigation.Top().Title, navigation.StackSummary())
	return tea.Batch(cmds...)
}

func (navigation NavigationStack) Top() NavigationItem {
	if len(navigation.stack) == 0 {
		return NavigationItem{}
	}
	return navigation.stack[len(navigation.stack)-1]
}

func (navigation *NavigationStack) Push(item NavigationItem) tea.Cmd {
	navigation.footer.UpdateContent(navigation.Top().Title, navigation.StackSummary())
	navigation.stack = append(navigation.stack, item)
	return navigation.initRoot()
}

func (navigation *NavigationStack) Replace(item NavigationItem) tea.Cmd {
	navigation.footer.UpdateContent(navigation.Top().Title, navigation.StackSummary())
	navigation.stack[len(navigation.stack)-1] = item
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
	navigation.footer.UpdateContent(navigation.Top().Title, navigation.StackSummary())
	cmds := []tea.Cmd{}
	var cmd tea.Cmd
	navigation.stack[len(navigation.stack)-1].Model, cmd = navigation.Top().Update(*navigation.size)
	cmds = append(cmds, cmd)
	navigation.stack[len(navigation.stack)-1].Model, cmd = navigation.Top().Update(ModelRestoreMsg{})
	cmds = append(cmds, cmd)
	navigation.footer.UpdateContent(navigation.Top().Title, navigation.StackSummary())
	return tea.Batch(cmds...)
}

func (navigation *NavigationStack) Size() tea.WindowSizeMsg {
	return *navigation.size
}
