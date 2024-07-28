package components

import (
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	tea "github.com/charmbracelet/bubbletea"
)

type Common struct {
	ID   string
	Size tea.WindowSizeMsg
}

type Interactive struct {
	Connection *database.ConnectionHandler
}