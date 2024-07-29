package taskform

import (
	"fmt"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	components.Common
	components.Interactive

	// Edit Fields
	Task task.Task

	// Form
	Form *huh.Form
	Data *FormData
}

type FormData struct {
	Code     string
	Title    string
	Content  string
	StatusId string
}

func CreateForm(model *Model) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Title:").
				Value(&model.Data.Title),

			huh.NewInput().
				Title("Code for testing, this should be auto generated:").
				Value(&model.Data.Code),

			huh.NewText().
				Value(&model.Data.Content).
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
		WithWidth(200).
		WithShowHelp(true).
		WithShowErrors(true).
		WithTheme(CustomStyles())
}

func CreateTaskForm() Model {
	model := Model{Data: &FormData{}}
	model.Form = CreateForm(&model)
	return model
}

func EditTaskForm(task task.Task) Model {
	model := Model{
		Task: task,
		Data: &FormData{Title: task.Name, Content: task.Content},
	}

	model.Form = CreateForm(&model)

	return model
}

func (m Model) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = msg
	}
}

func (m Model) View() string {
	return lipgloss.NewStyle().
		Padding(1, 4, 0, 1).Render(m.Form.View())
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
