package main

import (
	"fmt"
	"os"

	board "github.com/Dunkansdk/kanban-dunkan/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := board.NewKanban()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
