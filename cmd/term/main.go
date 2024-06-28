package main

import (
	"fmt"
	"os"

	board "github.com/Dunkansdk/kanban-dunkan/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()

	model := board.NewKanban()
	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
