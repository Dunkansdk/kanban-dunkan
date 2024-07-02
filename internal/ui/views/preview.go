package views

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	task task.Task
	size tea.WindowSizeMsg
}

func NewPreview(task task.Task) Model {
	return Model{task: task}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = msg
	}

	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("Task Code: %s\nTask Name: %s\nTask Description: %s", m.task.Code, m.task.Name, m.task.Content)
}
