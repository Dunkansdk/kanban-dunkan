package kanban

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/keyboard"
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/column"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/create"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/messages"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/preview"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Kanban struct {
	components.Common
	components.Interactive
	loaded       bool
	columns      []column.Model
	quitting     bool
	help         help.Model
	activeColumn *column.Model
}

var activeId int

func NewKanban(conn *database.ConnectionHandler) *Kanban {
	help := help.New()
	help.ShowAll = false
	kanban := Kanban{help: help}
	kanban.ID = zone.NewPrefix()
	kanban.Connection = conn
	return &kanban
}

func (kanban Kanban) Init() tea.Cmd {
	if !kanban.loaded {
		kanban.InitializeColumns()
		kanban.loaded = true
	}
	return nil
}

func (kanban Kanban) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	kanban.activeColumn, activeId = kanban.Active()

	var cmds []tea.Cmd

	switch message := message.(type) {
	case navigation.ModelRestoreMsg:
		return kanban, kanban.RefreshColumns()

	case tea.WindowSizeMsg:
		if !kanban.loaded {
			kanban.InitializeColumns()
			kanban.loaded = true
		}
		kanban.UpdateSize(message)
		for index, value := range kanban.columns {
			model, cmd := value.Update(message)
			kanban.columns[index] = model.(column.Model)
			cmds = append(cmds, cmd)
		}
		return kanban, tea.Batch(cmds...)

	case tea.KeyMsg:
		if !kanban.activeColumn.List.SettingFilter() {
			switch {
			case key.Matches(message, keyboard.Options.Quit):
				kanban.quitting = true
				return kanban, tea.Quit
			case key.Matches(message, keyboard.Options.Left):
				kanban.Prev()
			case key.Matches(message, keyboard.Options.Right):
				kanban.Next()
			case key.Matches(message, keyboard.Options.Refresh):
				return kanban, kanban.RefreshColumns()
			case key.Matches(message, keyboard.Options.Help):
				kanban.help.ShowAll = !kanban.help.ShowAll
			case key.Matches(message, keyboard.Options.Enter):
				view := preview.NewPreview(kanban.activeColumn, kanban.activeColumn.List.SelectedItem().(task.Task))
				return kanban, navigation.Push(navigation.NavigationItem{Title: "Preview", Model: view})
			case key.Matches(message, keyboard.Options.New):
				view := create.CreateTaskView(kanban.Connection)
				return kanban, navigation.Push(navigation.NavigationItem{Title: "Create task", Model: view})
			}
		}

	case tea.MouseMsg:
		switch message.Button {
		case tea.MouseButtonWheelUp:
			kanban.columns[activeId].List.CursorUp()
			return kanban, nil
		case tea.MouseButtonWheelDown:
			kanban.columns[activeId].List.CursorDown()
			return kanban, nil
		}
		if message.Action == tea.MouseActionPress || message.Button == tea.MouseButtonLeft {
			kanban.ZoneSelectLine(message)
		}
		if message.Action == tea.MouseActionMotion && !kanban.activeColumn.List.SettingFilter() {
			kanban.ZoneSelectColumn(message)
		}

	case messages.CreateTaskMsg:
		kanban.RefreshColumn(kanban.activeColumn, message.Task)
		return kanban, nil
	}

	model, cmd := kanban.activeColumn.Update(message)
	if _, ok := model.(column.Model); ok {
		kanban.columns[activeId] = model.(column.Model)
	} else {
		return model, cmd
	}

	return kanban, cmd
}

func (kanban Kanban) View() string {
	if kanban.quitting {
		return ""
	}
	if kanban.loaded {
		var components []string
		for _, column := range kanban.columns {
			components = append(components, column.View())
		}
		kanbanCmd := lipgloss.JoinHorizontal(
			lipgloss.Center,
			components...,
		)
		kanbanStyle := lipgloss.NewStyle().PaddingLeft(1).Render(kanbanCmd)
		return zone.Scan(kanbanStyle)
	} else {
		return "Loading\n"
	}
}
