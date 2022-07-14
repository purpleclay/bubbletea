package main

import (
	"log"

	"github.com/charmbracelet/bubbles/timer"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// this is an enum for Go
type sessionState uint

const (
	timerView sessionState = iota
	statsView
)

type mainModel struct {
	state sessionState
	timer timer.Model
	stats viewport.Model
}

func New() mainModel {
	// initialize your model; timerView is the first "view" we want to see
	return mainModel{state: timerView}
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle IO -> keypress, WindowSizeMSg
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
		if msg.String() == "s" {
			m.state = statsView
			return m, nil
		}
	case tea.WindowSizeMsg:
		// handle resizing windows
		// handle your Msgs
	}
	return m, nil
}

func (m mainModel) View() string {
	switch m.state {
	case statsView:
		return m.stats.View()
	default:
		return "timer is " + m.timer.View()
	}
}

func main() {
	p := tea.NewProgram(New())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
