package main

import (
	"image/color"

	"github.com/BigJk/crt"

	bubbleadapter "github.com/BigJk/crt/bubbletea"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

func main() {
	zone.NewGlobal()
	model := views.NewKanban()

	// Load fonts for normal, bold and italic text styles.
	fonts, err := crt.LoadFaces("./assets/fonts/IosevkaTermNerdFontMono-Regular.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Bold.ttf", "./assets/fonts/IosevkaTermNerdFontMono-Italic.ttf", crt.GetFontDPI(), 14.0)
	if err != nil {
		panic(err)
	}

	// Just pass your tea.Model to the bubbleadapter, and it will render it to the terminal.
	win, _, err := bubbleadapter.Window(1000, 600, fonts, model, color.Black, tea.WithAltScreen(), tea.WithMouseAllMotion())
	if err != nil {
		panic(err)
	}

	// Star the terminal with the given title.
	if err := win.Run("Kandun"); err != nil {
		panic(err)
	}
}
