package preview

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/column"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	column   *column.Model
	task     task.Task
	size     tea.WindowSizeMsg
	viewport viewport.Model
}

func NewPreview(column *column.Model, selected task.Task) Model {
	vp := viewport.New(0, 0)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("130")).
		PaddingRight(2)
	markdown, _ := glamour.Render(selected.Content, "dark")
	vp.SetContent(markdown)

	return Model{column: column, task: selected, viewport: vp}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = msg
		m.viewport.Height = msg.Height - 7
		m.viewport.Width = msg.Width - m.column.Size.Width - 4
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	titleContent := fmt.Sprintf("%s\n%s", m.task.Name, m.task.Code)
	title := lipgloss.NewStyle().
		Width(m.viewport.Width).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("130")).
		Padding(0, 2, 0, 1).
		Render(titleContent)

	taskStyle := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		m.viewport.View())

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.column.View(),
		taskStyle,
	)
}
