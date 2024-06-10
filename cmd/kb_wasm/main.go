//go:build js
// +build js

package main

import (
	"bytes"
	"fmt"
	"os"
	"syscall/js"
	"time"

	board "github.com/Dunkansdk/kanban-dunkan/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type MinReadBuffer struct {
	buf *bytes.Buffer
}

// For some reason bubbletea doesn't like a Reader that will return 0 bytes instead of blocking,
// so we use this hacky workaround for now. As javascript is single threaded this should be fine
// with regard to concurrency.
func (b *MinReadBuffer) Read(p []byte) (n int, err error) {
	for b.buf.Len() == 0 {
		time.Sleep(100 * time.Millisecond)
	}
	return b.buf.Read(p)
}

func (b *MinReadBuffer) Write(p []byte) (n int, err error) {
	return b.buf.Write(p)
}

// Creates the bubbletea program and registers the necessary functions in javascript
func createTeaForJS(model tea.Model, option ...tea.ProgramOption) *tea.Program {
	// Create buffers for input and output
	fromJs := &MinReadBuffer{buf: bytes.NewBuffer(nil)}
	fromGo := bytes.NewBuffer(nil)

	prog := tea.NewProgram(model, append([]tea.ProgramOption{tea.WithInput(fromJs), tea.WithOutput(fromGo)}, option...)...)

	// Register write function in WASM
	js.Global().Set("bubbletea_write", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fromJs.Write([]byte(args[0].String()))
		fmt.Println("Wrote to Go:", args[0].String())
		return nil
	}))

	// Register read function in WASM
	js.Global().Set("bubbletea_read", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		b := make([]byte, fromGo.Len())
		_, _ = fromGo.Read(b)
		fromGo.Reset()
		if len(b) > 0 {
			fmt.Println("Read from Go:", string(b))
		}
		return string(b)
	}))

	// Register resize function in WASM
	js.Global().Set("bubbletea_resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width := args[0].Int()
		height := args[1].Int()
		prog.Send(tea.WindowSizeMsg{Width: width, Height: height})
		return nil
	}))

	return prog
}

func main() {
	model := board.New()
	prog := createTeaForJS(model, tea.WithAltScreen())

	fmt.Println("Starting program...")
	if _, err := prog.Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
