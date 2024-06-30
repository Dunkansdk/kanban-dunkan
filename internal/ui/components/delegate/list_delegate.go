package delegate

import (
	"fmt"
	"io"
	"strings"

	"github.com/Dunkansdk/kanban-dunkan/internal/task"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/reflow/truncate"
)

const (
	bullet   = "•"
	ellipsis = "…"
)

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(3)
	selectedItemStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				PaddingLeft(2).Foreground(lipgloss.Color("166"))
)

type ListCustomDelegate struct{}

const OFFSET = 4

func (customDelegate ListCustomDelegate) Height() int                             { return 2 }
func (customDelegate ListCustomDelegate) Spacing() int                            { return 1 }
func (customDelegate ListCustomDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (customDelegate ListCustomDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, code string
	)

	if i, ok := item.(task.Task); ok {
		title = i.Title()
		code = i.Description()
	} else {
		return
	}

	textwidth := uint(m.Width() - itemStyle.GetPaddingLeft() - itemStyle.GetPaddingRight() - OFFSET)
	title = truncate.StringWithTail(title, textwidth, ellipsis)

	str := fmt.Sprintf("%s\n%s", title, code)

	fn := itemStyle.Render
	zone.Mark(code+title, itemStyle.Render(str))
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render(strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
