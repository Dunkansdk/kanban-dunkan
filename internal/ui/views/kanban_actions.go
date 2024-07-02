package views

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	tea "github.com/charmbracelet/bubbletea"
)

func (kanban *Kanban) Active() (*components.Column, int) {
	for index, column := range kanban.columns {
		if column.Focused() {
			return &kanban.columns[index], index
		}
	}
	return nil, 0
}

func (kanban *Kanban) RetreiveTasks() {
	taskRepository := task.NewTaskRepository()
	statuses := taskRepository.GetAllStatuses()

	kanban.columns = make([]components.Column, len(statuses))

	for index, value := range statuses {
		tasks, _ := taskRepository.GetAllByStatus(value)
		kanban.columns[value.ID].FillColumn(value, tasks)
		if index == 0 {
			kanban.columns[value.ID].Focus()
		}
	}
}

func (kanban *Kanban) UpdateSize(size tea.WindowSizeMsg) {
	for _, column := range kanban.columns {
		column.SetSize(size.Width, size.Height)
	}
}

func (kanban *Kanban) Next() {
	_, id := kanban.Active()
	if id < len(kanban.columns)-1 {
		kanban.columns[id].Blur()
		kanban.columns[id+1].Focus()
	}
}

func (kanban *Kanban) Prev() {
	_, id := kanban.Active()
	if id > 0 {
		kanban.columns[id].Blur()
		kanban.columns[id-1].Focus()
	}
}

func (kanban *Kanban) Select(index int) {
	kanban.activeColumn.Blur()
	kanban.columns[index].Focus()
}
