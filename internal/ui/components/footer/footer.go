package footer

import (
	"time"

	"github.com/Dunkansdk/kanban-dunkan/internal/ui/components"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/reflow/truncate"
)

// Height represents the height of the statusbar.
const height = 1

type Model struct {
	components.Common
	Mode      string
	UpdatedAt string
	Styles    FooterStyles
}

// ColorConfig
type ColorConfig struct {
	Foreground lipgloss.AdaptiveColor
	Background lipgloss.AdaptiveColor
}

type FooterStyles struct {
	FirstColumnColors  ColorConfig
	SecondColumnColors ColorConfig
	ThirdColumnColors  ColorConfig
	FourthColumnColors ColorConfig
}

// New creates a new instance of the UI.
func New(mode string) Model {
	footerStyles := FooterStyles{
		FirstColumnColors: ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "15", Light: "0"},
			Background: lipgloss.AdaptiveColor{Light: "208", Dark: "202"},
		},
		SecondColumnColors: ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "15", Light: "0"},
			Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		ThirdColumnColors: ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "15", Light: "0"},
			Background: lipgloss.AdaptiveColor{Light: "172", Dark: "166"},
		},
		FourthColumnColors: ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "15", Light: "0"},
			Background: lipgloss.AdaptiveColor{Light: "208", Dark: "202"},
		},
	}

	footer := Model{
		Mode:      mode,
		UpdatedAt: time.Now().Format(time.RFC822),
		Styles:    footerStyles,
	}

	footer.ID = zone.NewPrefix()

	return footer
}

// Init intializes the UI.
func (m Model) Init() tea.Cmd {
	return nil
}

// SetSize sets the width of the statusbar.
func (m *Model) SetSize(width int) {
	m.Size.Width = width
}

// Update updates the size of the statusbar.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width)
	}

	return m, nil
}

// View returns a string representation of a statusbar.
func (m Model) View() string {
	width := lipgloss.Width

	firstColumn := lipgloss.NewStyle().
		Foreground(m.Styles.FirstColumnColors.Foreground).
		Background(m.Styles.FirstColumnColors.Background).
		Padding(0, 1).
		Height(height).
		Render(truncate.StringWithTail(m.Mode, 30, "..."))

	thirdColumn := lipgloss.NewStyle().
		Foreground(m.Styles.ThirdColumnColors.Foreground).
		Background(m.Styles.ThirdColumnColors.Background).
		Align(lipgloss.Right).
		Padding(0, 1).
		Height(height).
		Render(m.UpdatedAt)

	fourthColumn := lipgloss.NewStyle().
		Foreground(m.Styles.FourthColumnColors.Foreground).
		Background(m.Styles.FourthColumnColors.Background).
		Padding(0, 1).
		Height(height).
		Render("KD")

	secondColumn := lipgloss.NewStyle().
		Foreground(m.Styles.SecondColumnColors.Foreground).
		Background(m.Styles.SecondColumnColors.Background).
		Padding(0, 1).
		Height(height).
		Width(m.Size.Width - width(firstColumn) - width(thirdColumn) - width(fourthColumn)).
		Render(truncate.StringWithTail(
			"we need a breadcrumb component to put it here, testing the trim",
			uint(m.Size.Width-width(firstColumn)-width(thirdColumn)-width(fourthColumn)-3),
			"..."),
		)

	return lipgloss.JoinHorizontal(lipgloss.Top,
		firstColumn,
		secondColumn,
		thirdColumn,
		fourthColumn,
	)
}
