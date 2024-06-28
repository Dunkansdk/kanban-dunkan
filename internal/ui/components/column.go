package components

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

type Column struct {
	Common
	focus  bool
	Status task.TaskStatus
	List   list.Model
}

const DIVISOR_OFFSET = 4

func (column *Column) FillColumn(status task.TaskStatus, tasks []task.Task) {
	column.List = list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	column.List.SetShowHelp(false)
	column.List.Title = status.Name
	column.List.Styles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("130")).
		Padding(0, 1)

	column.ID = zone.NewPrefix()

	var task_list []list.Item
	for _, element := range tasks {
		task_list = append(task_list, element)
	}
	column.List.SetItems(task_list)
	column.List.SetShowStatusBar(true)
}

func (column Column) Init() tea.Cmd {
	return nil
}

func (column Column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		column.SetSize(msg.Width, msg.Height)
		column.List.SetSize(msg.Width/DIVISOR_OFFSET, msg.Height/2)

	case tea.KeyMsg:
		if column.List.FilterState() == list.Filtering {
			break
		}

		return column, nil
	}

	column.List, cmd = column.List.Update(msg)
	return column, cmd
}

func (column Column) View() string {
	return zone.Mark(column.ID+column.Status.Name, column.getStyle().Render(column.List.View()))
}

func (column *Column) Blur() {
	column.focus = false
}

func (column *Column) Focus() {
	column.focus = true
}

func (column *Column) Focused() bool {
	return column.focus
}

func (column *Column) SetSize(width int, height int) {
	column.Width = width / DIVISOR_OFFSET
}

func (column *Column) getStyle() lipgloss.Style {
	if column.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Height(column.Height - DIVISOR_OFFSET).
			Width(column.Width)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Height(column.Height - DIVISOR_OFFSET).
		Width(column.Width)
}
