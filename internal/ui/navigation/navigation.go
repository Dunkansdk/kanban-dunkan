package navigation

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/footer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NavigationStack struct {
	// zones *zone.Manager
	stack  []tea.Model
	size   *tea.WindowSizeMsg
	footer footer.Model
}

func NewNavigation(root tea.Model) NavigationStack {
	navigation := NavigationStack{
		stack:  []tea.Model{root},
		footer: footer.New("Preview"),
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
		navigation.size.Height -= 1
		navigation.footer.Size = msg
	}
	navigation.footer.Update(msg)
	var cmd tea.Cmd
	navigation.stack[len(navigation.stack)-1], cmd = navigation.Top().Update(msg)
	return navigation, cmd
}

func (navigation NavigationStack) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, navigation.Top().View(), navigation.footer.View())
}

func (navigation NavigationStack) StackSummary() string {
	return "Kanban -> Testing breadcrumbs"
}
