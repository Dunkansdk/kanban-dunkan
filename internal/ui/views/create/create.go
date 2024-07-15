package create

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	size      tea.WindowSizeMsg
	viewport  viewport.Model
	textInput textinput.Model
}

func CreateTaskView() Model {
	input := textinput.New()
	input.Placeholder = "Testing task"
	input.Focus()
	input.Width = 30
	return Model{textInput: input}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = msg
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func RenderInput(label string, input textinput.Model) string {
	return fmt.Sprintf(
		"%s\n%s",
		label,
		input.View(),
	) + "\n"
}

func (m Model) View() string {
	return RenderInput("Task Name: ", m.textInput)
}
