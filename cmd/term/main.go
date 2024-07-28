package main

import (
	"fmt"
	"os"

	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/kanban"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kandun",
	Short: "A CLI project management tool",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var kanbanCmd = &cobra.Command{
	Use:   "kanban",
	Short: "Interact with your tasks in a Kanban board.",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		zone.NewGlobal()

		// Setting up connection handler
		connectionHandler := database.CreateConnection(&database.SQLite3DB{})
		model := kanban.NewKanban(connectionHandler)

		navigation := navigation.NewNavigation("Board", model)
		p := tea.NewProgram(navigation, tea.WithAltScreen(), tea.WithMouseAllMotion())
		_, err := p.Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(kanbanCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
