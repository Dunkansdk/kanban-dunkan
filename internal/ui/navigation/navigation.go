package navigation

import (
	"strings"

	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/footer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NavigationStack struct {
	// zones *zone.Manager
	stack  []NavigationItem
	size   *tea.WindowSizeMsg
	footer *footer.Model
}

func NewNavigation(title string, model tea.Model) NavigationStack {
	navigation := NavigationStack{
		stack:  []NavigationItem{NavigationItem{Title: title, Model: model}},
		footer: footer.New("Preview"),
	}
	return navigation
}

func (navigation NavigationStack) Init() tea.Cmd {
	return navigation.initRoot()
}

func (navigation NavigationStack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	active := navigation.Top()
	switch msg := msg.(type) {
	case ModelPopMsg:
		return navigation, navigation.Pop()
	case ModelPushMsg:
		return navigation, navigation.Push(msg.Item)
	case ModelReplaceMsg:
		return navigation, navigation.Replace(msg.Item)
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return navigation, navigation.Pop()
		}
	case tea.WindowSizeMsg:
		navigation.size = &msg
		navigation.size.Height -= 1
		navigation.footer.Size = msg
	}
	navigation.footer.Update(msg)
	model, cmd := active.Update(msg)
	navigation.stack[len(navigation.stack)-1] = model.(NavigationItem)
	return navigation, cmd
}

func (navigation NavigationStack) View() string {
	top := navigation.Top()
	if top.Model == nil {
		return ""
	}
	return lipgloss.JoinVertical(lipgloss.Left, top.View(), navigation.footer.View())
}

func (navigation NavigationStack) StackSummary() string {
	var breadcrumb strings.Builder
	for index, item := range navigation.stack {
		breadcrumb.WriteString(item.Title)
		if index != len(navigation.stack)-1 {
			breadcrumb.WriteString(" -> ")
		}
	}
	return breadcrumb.String()
}
