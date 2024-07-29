package create

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/taskform"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/messages"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type Model struct {
	components.Common
	components.Interactive

	size tea.WindowSizeMsg
	Task task.Task

	// Form
	TaskForm taskform.Model
}

func CreateTaskView(conn *database.ConnectionHandler) Model {
	model := Model{
		TaskForm: taskform.CreateTaskForm(),
	}
	model.Connection = conn
	return model
}

func (m Model) Init() tea.Cmd {
	return m.TaskForm.Form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, navigation.Pop()
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.TaskForm.Form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.TaskForm.Form = f
		cmds = append(cmds, cmd)
	}

	if m.TaskForm.Form.State == huh.StateCompleted {
		return m, tea.Batch(messages.Create(&task.Task{
			Code:    m.TaskForm.Data.Code,
			Name:    m.TaskForm.Data.Title,
			Content: m.TaskForm.Data.Content,
		}), navigation.Pop())
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.TaskForm.View()
}
