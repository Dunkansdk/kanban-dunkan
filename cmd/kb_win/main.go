package main

import (
	"image/color"

	"github.com/BigJk/crt"

	bubbleadapter "github.com/BigJk/crt/bubbletea"
	board "github.com/Dunkansdk/kanban-dunkan/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := board.New()

	// Load fonts for normal, bold and italic text styles.
	fonts, err := crt.LoadFaces("./assets/fonts/IosevkaTermNerdFontMono-Regular.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Bold.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Italic.ttf", crt.GetFontDPI(), 14.0)
	if err != nil {
		panic(err)
	}

	// Just pass your tea.Model to the bubbleadapter, and it will render it to the terminal.
	win, _, err := bubbleadapter.Window(900, 600, fonts, model, color.Black, tea.WithAltScreen())
	if err != nil {
		panic(err)
	}

	// Star the terminal with the given title.
	if err := win.Run("Kandun"); err != nil {
		panic(err)
	}
}