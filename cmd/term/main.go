package main

import (
	"fmt"
	"os"

	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/kanban"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()

	model := kanban.NewKanban()
	navigation := navigation.NewNavigation(model)

	p := tea.NewProgram(navigation, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
