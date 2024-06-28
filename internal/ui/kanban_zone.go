package ui

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
		if len(item.Content) > 50 {
			if zone.Get(item.Code+item.Name).InBounds(message) || zone.Get(item.Code+item.Content[0:50]).InBounds(message) {
				kanban.activeColumn.List.Select(i)
			}
		} else {
			if zone.Get(item.Code+item.Name).InBounds(message) || zone.Get(item.Code+item.Content).InBounds(message) {
				kanban.activeColumn.List.Select(i)
			}
		}
	}
}
