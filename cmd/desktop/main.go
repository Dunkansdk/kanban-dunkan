package main

import (
	"fmt"
	"image/color"

	"github.com/BigJk/crt"

	bubbleadapter "github.com/BigJk/crt/bubbletea"
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/kanban"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()

	// Setting up connection handler
	connectionHandler := database.CreateConnection(&database.SQLite3DB{})
	model := kanban.NewKanban(connectionHandler)

	navigation := navigation.NewNavigation("Board", model)

	// Load fonts for normal, bold and italic text styles.
	fonts, err := crt.LoadFaces("./assets/fonts/IosevkaTermNerdFontMono-Regular.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Bold.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Italic.ttf", crt.GetFontDPI(), 14.0)
	if err != nil {
		panic(err)
	}

	// Just pass your tea.Model to the bubbleadapter, and it will render it to the terminal.
	win, _, err := Window(1000, 600, fonts, navigation, color.Black, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if err != nil {
		panic(err)
	}

	// Star the terminal with the given title.
	if err := win.Run("Kandun"); err != nil {
		panic(err)
	}
}

// Window creates a new crt based bubbletea window with the given width, height, fonts, model and default background color.
// Additional options can be passed to the bubbletea program.
func Window(width int, height int, fonts crt.Fonts, model tea.Model, defaultBg color.Color, options ...tea.ProgramOption) (*crt.Window, *tea.Program, error) {
	gameInput := crt.NewConcurrentRW()
	gameOutput := crt.NewConcurrentRW()

	go gameInput.Run()
	go gameOutput.Run()

	prog := tea.NewProgram(
		model,
		append([]tea.ProgramOption{
			tea.WithMouseAllMotion(),
			tea.WithInput(gameInput),
			tea.WithOutput(gameOutput),
		}, options...)...,
	)

	go func() {
		if _, err := prog.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
		}

		crt.SysKill()
	}()

	win, err := crt.NewGame(width, height, fonts, gameOutput, bubbleadapter.NewAdapter(prog), defaultBg)
	return win, prog, err
}
