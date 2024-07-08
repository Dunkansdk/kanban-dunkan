package kanban

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func (kanban *Kanban) ZoneSelectColumn(message tea.MouseMsg) {
	for index, column := range kanban.columns {
		if zone.Get(column.ID + column.Status.Name).InBounds(message) {
			kanban.activeColumn.Blur()
			kanban.columns[index].Focus()
		}
	}
}

func (kanban *Kanban) ZoneSelectLine(message tea.MouseMsg) {
	for i, listItem := range kanban.activeColumn.List.VisibleItems() {
		item, _ := listItem.(task.Task)
		if zone.Get(item.Code+item.Name).InBounds(message) || zone.Get(item.Code).InBounds(message) {
			kanban.activeColumn.List.Select(i)
		}
	}
}
