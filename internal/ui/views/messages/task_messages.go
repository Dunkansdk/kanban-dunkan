package messages

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	tea "github.com/charmbracelet/bubbletea"
)

type CreateTaskMsg struct {
	Task *task.Task
}

type UpdateMsg struct {
	Task *task.Task
}

func Create(task *task.Task) tea.Cmd {
	return tea.Batch(func() tea.Msg {
		return CreateTaskMsg{task}
	}, Update(task))
}

type EditTaskMsg struct {
	Task *task.Task
}

func Edit(task *task.Task) tea.Cmd {
	return func() tea.Msg {
		return EditTaskMsg{task}
	}
}

func Update(task *task.Task) tea.Cmd {
	return func() tea.Msg {
		return UpdateMsg{task}
	}
}
