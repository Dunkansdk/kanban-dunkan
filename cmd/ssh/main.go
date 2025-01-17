package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/navigation"
	"github.com/Dunkansdk/kanban-dunkan/internal/ui/views/kanban"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	zone "github.com/lrstanley/bubblezone"
	"github.com/muesli/termenv"
	"github.com/spf13/pflag"
)

// TODO: Configuration (.properties)
var (
	host = pflag.String("host", "localhost", "Address to listen to")
	port = pflag.Int("port", 42069, "Port to listen on")
)

func main() {
	pflag.Parse()
	run(*host, strconv.Itoa(*port))
}

func teaHandler(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := sess.Pty()
	if !active {
		log.Printf("no active terminal, skipping")
		return nil, nil
	}

	lipgloss.SetColorProfile(termenv.TrueColor)
	zone.NewGlobal()
	connectionHandler := database.CreateConnection(&database.PSQLDB{})
	model := kanban.NewKanban(connectionHandler)
	navigation := navigation.NewNavigation("Board", model)
	return navigation, []tea.ProgramOption{
		tea.WithMouseAllMotion(),
		tea.WithAltScreen(),
	}
}

func run(host string, port string) {
	/// Wish ssh application
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			activeterm.Middleware(), // Bubble Tea apps usually require a PTY.
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}
