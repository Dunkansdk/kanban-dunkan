package column

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/column/delegate"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

const CENTER_FACTOR = 5

type Model struct {
	components.Common
	focus   bool
	Status  task.TaskStatus
	List    list.Model
	divisor int
}

func (column *Model) FillColumn(status task.TaskStatus, tasks []task.Task, totalColumns int) {
	column.ID = zone.NewPrefix()
	column.divisor = totalColumns

	column.List = list.New([]list.Item{}, delegate.ListCustomDelegate{}, 0, 0)
	column.List.SetShowHelp(false)
	column.List.Title = status.Name
	column.List.Styles.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("130")).
		Padding(0, 1)

	var task_list []list.Item
	for _, element := range tasks {
		task_list = append(task_list, element)
	}
	column.List.SetItems(task_list)

	column.List.KeyMap.NextPage.SetEnabled(false)
	column.List.KeyMap.PrevPage.SetEnabled(false)
	column.List.KeyMap.GoToStart.SetEnabled(false)
	column.List.KeyMap.GoToEnd.SetEnabled(false)
}

func (column Model) Init() tea.Cmd {
	return nil
}

func (column Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		column.SetSize(msg.Width, msg.Height)
		column.List.SetSize((msg.Width/column.divisor)-CENTER_FACTOR, msg.Height-CENTER_FACTOR)

	case tea.KeyMsg:
		if column.List.FilterState() == list.Filtering {
			break
		}
	}

	column.List, cmd = column.List.Update(msg)
	return column, cmd
}

func (column Model) View() string {
	return zone.Mark(column.ID+column.Status.Name, column.getStyle().Render(column.List.View()))
}

func (column *Model) Blur() {
	column.focus = false
}

func (column *Model) Focus() {
	column.focus = true
}

func (column *Model) Focused() bool {
	return column.focus
}

func (column *Model) SetSize(width int, height int) {
	column.Size.Width = (width / column.divisor) - CENTER_FACTOR
}

func (column *Model) getStyle() lipgloss.Style {
	if column.Focused() {
		return lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("130")).
			Height(column.Size.Height - column.divisor).
			Width(column.Size.Width)
	}
	return lipgloss.NewStyle().
		Padding(1, 2).
		Border(lipgloss.HiddenBorder()).
		Height(column.Size.Height - column.divisor).
		Width(column.Size.Width)
}
