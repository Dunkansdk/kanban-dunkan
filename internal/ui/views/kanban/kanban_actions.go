package kanban

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/column"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/footer"
	tea "github.com/charmbracelet/bubbletea"
)

func (kanban *Kanban) Active() (*column.Model, int) {
	for index, column := range kanban.columns {
		if column.Focused() {
			return &kanban.columns[index], index
		}
	}
	return nil, 0
}

func (kanban *Kanban) InitializeColumns() {
	taskRepository := task.NewTaskRepository(kanban.Connection)
	statuses := taskRepository.GetAllStatuses()

	kanban.columns = make([]column.Model, len(statuses))

	for index, value := range statuses {
		tasks, _ := taskRepository.GetAllByStatus(value)
		kanban.columns[value.ID].CreateColumns(value, tasks, len(statuses))
		if index == 0 {
			kanban.columns[value.ID].Focus()
		}
	}
}

func (kanban *Kanban) RefreshColumn(active *column.Model, item *task.Task) {
	taskRepository := task.NewTaskRepository(kanban.Connection)
	taskRepository.Insert(item)
	tasks, _ := taskRepository.GetAllByStatus(active.Status)
	active.Refresh(tasks)
}

func (kanban *Kanban) RefreshColumns() tea.Cmd {
	taskRepository := task.NewTaskRepository(kanban.Connection)
	statuses := taskRepository.GetAllStatuses()
	for _, value := range statuses {
		tasks, _ := taskRepository.GetAllByStatus(value)
		kanban.columns[value.ID].Refresh(tasks)
	}
	// log.Info("Refreshing board")
	return func() tea.Msg {
		return footer.RefreshLastUpdated{}
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
