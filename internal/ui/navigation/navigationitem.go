package navigation

import tea "github.com/charmbracelet/bubbletea"

type NavigationItem struct {
	Title string
	Model tea.Model
}

func (item NavigationItem) Init() tea.Cmd {
	return item.Model.Init()
}

func (item NavigationItem) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := item.Model.Update(message)
	item.Model = model
	return item, cmd
}

func (item NavigationItem) View() string {
	return item.Model.View()
}
