package create

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/taskform"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	components.Common

	size tea.WindowSizeMsg
	Task task.Task

	// Form
	TaskForm taskform.Model
}

func CreateTaskView() Model {
	return Model{
		TaskForm: taskform.CreateTaskForm(),
	}
}

func (m Model) Init() tea.Cmd {
	return m.TaskForm.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg
	}

	return m.TaskForm.Update(msg)
}

func (m Model) View() string {
	return m.TaskForm.View()
}
