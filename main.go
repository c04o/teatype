package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// holds the program state
type model struct {
	targetText string
	userInput  string
}

// test text
func initialModel() model {
	return model{
		targetText: "the quick brown fox jumps over the lazy dog",
		// init blank
		userInput: "",
	}
}

// init i/o (like timers)
func (m model) Init() tea.Cmd {
	return nil
}

// handle events (key presses, etc.)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// quit
		case "ctrl+c", "esc":
			return m, tea.Quit

		// if backspace
		case "backspace":
			if len(m.userInput) > 0 {
				// reduce input
				m.userInput = m.userInput[:len(m.userInput)-1]
			}
		default:
			// handle standard character keys for now
			if len(msg.String()) == 1 {
				m.userInput += msg.String()
			}
		}
	}
	return m, nil
}

// render ui
func (m model) View() string {
	s := "Type the following:\n\n"
	s += m.targetText + "\n\n"
	s += "Your input: " + m.userInput + "\n\n"
	s += "(press esc to quit)"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
