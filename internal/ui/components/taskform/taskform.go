package taskform

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	components.Common

	// Edit Fields
	Task task.Task

	// Form
	form *huh.Form
	data *FormData
}

type FormData struct {
	Title    string
	Content  string
	StatusId string
}

func CreateForm(model *Model) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Title:").
				Value(&model.data.Title),

			huh.NewText().
				Value(&model.data.Content).
				Title("Content").
				Editor("vim").
				Description("Give a small description"),

			huh.NewSelect[string]().
				Key("status").
				Options(huh.NewOptions("To do", "In progress", "In requirements")...).
				Title("Choose status").
				Description("This will determine the status of the task"),

			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("Welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(100).
		WithShowHelp(true).
		WithShowErrors(true).
		WithTheme(CustomStyles())
}

func CreateTaskForm() Model {
	model := Model{}
	model.form = CreateForm(&model)
	return model
}

func EditTaskForm(task task.Task) Model {
	model := Model{
		Task: task,
	}

	model.data = &FormData{Title: task.Name, Content: task.Content}
	model.form = CreateForm(&model)

	return model
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, navigation.Pop()
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Padding(1, 4, 0, 1).Render(m.form.View())
}

func CustomStyles() *huh.Theme {
	t := huh.ThemeBase()

	var (
		background = lipgloss.AdaptiveColor{Dark: "235"}
		selection  = lipgloss.AdaptiveColor{Dark: "166"}
		foreground = lipgloss.AdaptiveColor{Dark: "223"}
		comment    = lipgloss.AdaptiveColor{Dark: "245"}
		green      = lipgloss.AdaptiveColor{Dark: "108"}
		orange     = lipgloss.AdaptiveColor{Dark: "166"}
		red        = lipgloss.AdaptiveColor{Dark: "124"}
		yellow     = lipgloss.AdaptiveColor{Dark: "172"}
	)

	t.Focused.Base = t.Focused.Base.BorderForeground(selection)
	t.Focused.Title = t.Focused.Title.Foreground(orange)
	t.Focused.NoteTitle = t.Focused.NoteTitle.Foreground(orange)
	t.Focused.Description = t.Focused.Description.Foreground(comment)
	t.Focused.ErrorIndicator = t.Focused.ErrorIndicator.Foreground(red)
	t.Focused.Directory = t.Focused.Directory.Foreground(orange)
	t.Focused.File = t.Focused.File.Foreground(foreground)
	t.Focused.ErrorMessage = t.Focused.ErrorMessage.Foreground(red)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(yellow)
	t.Focused.NextIndicator = t.Focused.NextIndicator.Foreground(yellow)
	t.Focused.PrevIndicator = t.Focused.PrevIndicator.Foreground(yellow)
	t.Focused.Option = t.Focused.Option.Foreground(foreground)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(yellow)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(green)
	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.Foreground(green)
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(foreground)
	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.Foreground(comment)
	t.Focused.FocusedButton = t.Focused.FocusedButton.Foreground(foreground).Background(orange).Bold(true)
	t.Focused.BlurredButton = t.Focused.BlurredButton.Foreground(foreground).Background(background)

	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(yellow)
	t.Focused.TextInput.Placeholder = t.Focused.TextInput.Placeholder.Foreground(comment)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(yellow)

	t.Blurred = t.Focused
	t.Blurred.Base = t.Blurred.Base.BorderStyle(lipgloss.HiddenBorder())
	t.Blurred.NextIndicator = lipgloss.NewStyle()
	t.Blurred.PrevIndicator = lipgloss.NewStyle()

	return t
}
