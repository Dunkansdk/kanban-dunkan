package main

import (
	"fmt"
	"os"

	"github.com/Dunkansdk/kanban-dunkan/internal/ui"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()

	model := views.NewKanban()
	navigation := ui.NewNavigation(model)

	p := tea.NewProgram(navigation, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
