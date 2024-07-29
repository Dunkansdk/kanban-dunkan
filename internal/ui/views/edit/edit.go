package edit

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/taskform"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	components.Common
	components.Interactive

	size tea.WindowSizeMsg
	Task task.Task

	// Form
	TaskForm taskform.Model
}

func EditTaskView(conn *database.ConnectionHandler, task task.Task) Model {
	model := Model{
		Task:     task,
		TaskForm: taskform.EditTaskForm(task),
	}
	model.Connection = conn
	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg
	}

	return m, nil
}

func (m Model) View() string {
	return m.TaskForm.View()
}
