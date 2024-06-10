package board

import (
	"strings"
	"time"

	TaskImpl "github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const DIVISOR_OFFSET = 4

type boardStyles struct {
	column  lipgloss.Style
	focused lipgloss.Style
	help    lipgloss.Style
	title   lipgloss.Style
}

type tickMsg time.Time

type Model struct {
	loaded   bool
	focused  TaskImpl.TaskStatus
	lists    []list.Model
	quitting bool
	columns  []TaskImpl.TaskStatus
	styles   boardStyles
	progress progress.Model
}

func makeStyles(renderer *lipgloss.Renderer) boardStyles {
	return boardStyles{
		column: renderer.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder()),
		focused: renderer.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("166")),
		help: renderer.NewStyle().
			Foreground(lipgloss.Color("242")),
		title: renderer.NewStyle().
			Background(lipgloss.Color("130")).
			Padding(0, 1),
	}
}

func New() *Model {
	renderer := lipgloss.DefaultRenderer()
	return &Model{
		// progress bar
		styles:   makeStyles(renderer),
		progress: progress.New(progress.WithColorProfile(renderer.ColorProfile()), progress.WithSolidFill("130")),
	}
}

func NewWithRenderer(renderer *lipgloss.Renderer) *Model {
	return &Model{
		// progress bar
		styles:   makeStyles(renderer),
		progress: progress.New(progress.WithColorProfile(renderer.ColorProfile()), progress.WithSolidFill("130")),
	}
}

func (model *Model) Next() {
	if model.focused == model.columns[len(model.columns)-1] {
		model.focused = model.columns[0]
	} else {
		model.focused = model.columns[model.focused.ID+1] // Esto se va a romper espectacularmente.
	}
}

func (model *Model) Prev() {
	if model.focused == model.columns[0] {
		model.focused = model.columns[len(model.columns)-1]
	} else {
		model.focused = model.columns[model.focused.ID-1]
	}
}

func (model *Model) initBoard(width, height int) {

	// TODO: I should divide this by project.

	taskRepository := TaskImpl.NewTaskRepository()

	// Init TaskStatus
	model.columns = taskRepository.GetAllStatuses()

	default_list := list.New([]list.Item{}, list.NewDefaultDelegate(), width/DIVISOR_OFFSET, height/2)
	default_list.SetShowHelp(false)
	default_list.SetDelegate(list.NewDefaultDelegate())
	default_list.SetSpinner(spinner.Ellipsis)

	model.lists = make([]list.Model, len(model.columns))

	for _, value := range model.columns {
		model.lists[value.ID] = default_list
	}

	for _, value := range model.columns {
		model.lists[value.ID].Title = value.Name
		model.lists[value.ID].Styles.Title = model.styles.title
		tasks, _ := taskRepository.GetAllByStatus(value)
		var task_list []list.Item
		for _, element := range tasks {
			task_list = append(task_list, &element)
		}
		model.lists[value.ID].SetItems(task_list)
		model.lists[value.ID].SetShowStatusBar(true)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (model Model) Init() tea.Cmd {
	return tickCmd()
}

func (model Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		// kanban
		if !model.loaded {
			model.styles.column.Width(message.Width / DIVISOR_OFFSET)
			model.styles.column.Height(message.Height - DIVISOR_OFFSET)
			model.styles.focused.Width(message.Width / DIVISOR_OFFSET)
			model.styles.focused.Height(message.Height - DIVISOR_OFFSET)
			model.initBoard(message.Width, message.Height)

			// progress bar
			model.progress.Width = message.Width
			if model.progress.Width > 80 {
				model.progress.Width = 80
			}
		}

	case tea.KeyMsg:
		switch message.String() {
		case "ctrl+c", "q":
			model.quitting = true
			return model, tea.Quit
		case "left", "h":
			model.Prev()
		case "right", "l":
			model.Next()
		}

	// progress bar
	case tickMsg:
		if !model.loaded {
			if model.progress.Percent() == 1.0 {
				model.loaded = true
			}
			cmd := model.progress.IncrPercent(0.40)
			return model, tea.Batch(tickCmd(), cmd)
		}

	case progress.FrameMsg:
		if !model.loaded {
			progressModel, cmd := model.progress.Update(message)
			model.progress = progressModel.(progress.Model)
			return model, cmd
		}
	}

	var cmd tea.Cmd
	model.lists[model.focused.ID], cmd = model.lists[model.focused.ID].Update(message)
	return model, cmd
}

func (model Model) View() string {
	if model.quitting {
		return ""
	}
	if model.loaded {
		// This should be defined in the board_impl.go?????
		var c_styles []string
		for _, value := range model.columns {
			if value == model.focused || (model.focused.Name == "" && value.ID == 0) {
				c_styles = append(c_styles, model.styles.focused.Render(model.lists[value.ID].View()))
			} else {
				c_styles = append(c_styles, model.styles.column.Render(model.lists[value.ID].View()))
			}
		}
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			c_styles...,
		)
	} else {
		// progress bar
		pad := strings.Repeat(" ", 2)
		return "\n" +
			pad + model.progress.View() + "\n\n" +
			pad + model.styles.help.Render("Press any key to quit")
	}
}
