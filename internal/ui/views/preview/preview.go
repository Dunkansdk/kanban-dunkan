package preview

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	components.Interactive
	components.Common

	task     task.Task
	viewport viewport.Model
}

func NewPreview(selected task.Task) Model {
	vp := viewport.New(0, 0)
	vp.Style = lipgloss.NewStyle().PaddingRight(2)
	markdown, _ := glamour.Render(selected.Content, "dark")
	vp.SetContent(markdown)
	vp.Height = 15
	vp.Width = 30

	return Model{task: selected, viewport: vp}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder(), true).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("130")).
		Padding(1, 2).
		Render(taskStyle)
}
