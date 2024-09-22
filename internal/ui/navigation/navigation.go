package navigation

import (
	"strings"

	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/footer"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components/overlay"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NavigationStack struct {
	// zones    *zone.Manager
	stack    []NavigationItem
	size     *tea.WindowSizeMsg
	footer   *footer.Model
	tooltips map[string]Tooltip
	viewport viewport.Model
}

func NewNavigation(title string, model tea.Model) NavigationStack {
	navigation := NavigationStack{
		stack:    []NavigationItem{NavigationItem{Title: title, Model: model}},
		footer:   footer.New("Preview"),
		viewport: viewport.New(0, 0),
		tooltips: map[string]Tooltip{},
	}
	return navigation
}

func (navigation NavigationStack) Init() tea.Cmd {
	return navigation.initRoot()
}

func (navigation NavigationStack) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	active := navigation.Top()
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		navigation.size = &msg
		navigation.footer.Size = msg
		navigation.viewport.Height = msg.Height - footer.Height
		navigation.viewport.Width = msg.Width
	case TooltipMsg:
		navigation.tooltips[msg.ID] = Tooltip(msg)
	case TooltipDeleteMsg:
		delete(navigation.tooltips, string(msg))
	case TooltipClearMsg:
		navigation.tooltips = map[string]Tooltip{}
	case ModelPopMsg:
		return navigation, navigation.Pop()
	case ModelPushMsg:
		return navigation, navigation.Push(msg.Item)
	case ModelReplaceMsg:
		return navigation, navigation.Replace(msg.Item)
	case footer.RefreshLastUpdated:
		model, command := navigation.footer.Update(msg)
		navigation.footer = &model
		return navigation, command
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if len(navigation.tooltips) != 0 {
				navigation.tooltips = map[string]Tooltip{}
				return navigation, nil
			} else {
				return navigation, navigation.Pop()
			}
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Footer
	_, cmd = navigation.footer.Update(msg)
	cmds = append(cmds, cmd)

	// Active Navigation
	model, activecmd := active.Update(msg)
	navigation.stack[len(navigation.stack)-1] = model.(NavigationItem)
	cmds = append(cmds, activecmd)

	// Viewport
	navigation.viewport, cmd = navigation.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return navigation, tea.Batch(cmds...)
}

func (navigation NavigationStack) View() string {
	top := navigation.Top()
	if top.Model == nil {
		return ""
	}
	navigation.viewport.SetContent(top.View())
	view := lipgloss.JoinVertical(lipgloss.Left, navigation.viewport.View(), navigation.footer.View())

	for _, v := range navigation.tooltips {
		view = overlay.PlaceOverlay(v.X, v.Y, v.Content, view)
	}

	return view
}

func (navigation NavigationStack) StackSummary() string {
	var breadcrumb strings.Builder
	for index, item := range navigation.stack {
		breadcrumb.WriteString(item.Title)
		if index != len(navigation.stack)-1 {
			breadcrumb.WriteString(" ➤  ")
		}
	}
	return breadcrumb.String()
}
