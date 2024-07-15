package taskform

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	components.Common

	// Edit Fields
	Task task.Task

	// Form
	NameInput    textinput.Model
	ContentInput textarea.Model
}

func EditTaskForm(task task.Task) Model {
	return Model{
		Task:         task,
		NameInput:    createInput(task.Name, 50, true),
		ContentInput: createArea(task.Content),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg
		m.ContentInput.SetWidth(msg.Width)
		m.ContentInput.SetHeight(msg.Height - 10)

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.NameInput.Focused() {
				m.NameInput.Blur()
				m.ContentInput.Focus()
				m.ContentInput.SetCursor(0)
				return m, textarea.Blink
			}
		}
	}

	if m.NameInput.Focused() {
		m.NameInput, cmd = m.NameInput.Update(msg)
	} else {
		m.ContentInput, cmd = m.ContentInput.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.NameInput.View(),
		m.ContentInput.View(),
		"(esc to quit)",
	) + "\n"
}

func createInput(placeholder string, charlimit int, focus bool) textinput.Model {
	input := textinput.New()
	input.Placeholder = placeholder
	if focus {
		input.Focus()
	}
	input.CharLimit = charlimit
	return input
}

func createArea(placeholder string) textarea.Model {
	area := textarea.New()
	area.Prompt = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).PaddingLeft(2).Render("┃ ")
	area.CharLimit = 0
	area.SetWidth(30)
	area.SetValue(placeholder)
	return area
}
