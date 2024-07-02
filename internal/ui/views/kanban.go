package views

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/keyboard"
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Kanban struct {
	components.Common
	loaded       bool
	columns      []components.Column
	quitting     bool
	help         help.Model
	activeColumn *components.Column
	motion       bool
}

var activeId int

func NewKanban() *Kanban {
	help := help.New()
	help.ShowAll = false
	kanban := Kanban{motion: true, help: help}
	kanban.ID = zone.NewPrefix()
	return &kanban
}

func (kanban Kanban) Init() tea.Cmd {
	if !kanban.loaded {
		kanban.RetreiveTasks()
		kanban.loaded = true
	}
	return nil
}

func (kanban Kanban) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	kanban.activeColumn, activeId = kanban.Active()

	var cmds []tea.Cmd

	switch message := message.(type) {
	case ui.ModelRestoreMsg:
		for index, column := range kanban.columns {
			model, cmd := column.Update(message)
			kanban.columns[index] = model.(components.Column)
			cmds = append(cmds, cmd)
		}
		return kanban, tea.Batch(cmds...)

	case tea.WindowSizeMsg:
		if !kanban.loaded {
			kanban.RetreiveTasks()
			kanban.loaded = true
		}
		kanban.UpdateSize(message)
		for index, column := range kanban.columns {
			model, cmd := column.Update(message)
			kanban.columns[index] = model.(components.Column)
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
			case key.Matches(message, keyboard.Options.Help):
				kanban.help.ShowAll = !kanban.help.ShowAll
			case key.Matches(message, keyboard.Options.Enter):
				return kanban, ui.Push(NewPreview(kanban.activeColumn.List.SelectedItem().(task.Task)))
			}
		}

	case tea.MouseMsg:
		if kanban.motion {
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
		}
	}

	model, cmd := kanban.activeColumn.Update(message)
	if _, ok := model.(components.Column); ok {
		kanban.columns[activeId] = model.(components.Column)
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
		var c_styles []string
		for _, column := range kanban.columns {
			c_styles = append(c_styles, column.View())
		}
		kanbanStyle := lipgloss.JoinHorizontal(
			lipgloss.Left,
			c_styles...,
		)

		return zone.Scan(lipgloss.JoinVertical(lipgloss.Left, kanbanStyle, kanban.help.View(keyboard.Options)))
	} else {
		return "Loading\n"
	}
}
