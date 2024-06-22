package ui

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/keyboard"
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Kanban struct {
	loaded   bool
	columns  []components.Column
	quitting bool
	help     help.Model
}

func NewKanban() *Kanban {
	help := help.New()
	help.ShowAll = false
	return &Kanban{help: help}
}

func (kanban *Kanban) RetreiveTasks(width, height int) {
	// TODO: I should divide this by project.
	taskRepository := task.NewTaskRepository()
	statuses := taskRepository.GetAllStatuses()

	kanban.columns = make([]components.Column, len(statuses))

	for index, value := range statuses {
		tasks, _ := taskRepository.GetAllByStatus(value)
		kanban.columns[value.ID].FillColumn(value, tasks)
		kanban.columns[value.ID].SetSize(width, height)
		if index == 0 {
			kanban.columns[value.ID].Focus()
		}
	}
}

func (kanban Kanban) Init() tea.Cmd {
	return nil
}

func (kanban *Kanban) Active() (*components.Column, int) {
	for index, column := range kanban.columns {
		if column.Focused() {
			return &column, index
		}
	}
	return nil, 0
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

func (kanban Kanban) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		var cmds []tea.Cmd
		kanban.help.Width = message.Width - 4
		if !kanban.loaded {
			kanban.RetreiveTasks(message.Width, message.Height)
		}
		for index, column := range kanban.columns {
			model, cmd := column.Update(message)
			kanban.columns[index] = model.(components.Column)
			cmds = append(cmds, cmd)
		}
		kanban.loaded = true
		return kanban, tea.Batch(cmds...)

	case tea.KeyMsg:
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
		}
	}

	_, activeId := kanban.Active()
	model, cmd := kanban.columns[activeId].Update(message)
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

		return lipgloss.JoinVertical(lipgloss.Left, kanbanStyle, kanban.help.View(keyboard.Options))
	} else {
		return "Loading\n"
	}
}
