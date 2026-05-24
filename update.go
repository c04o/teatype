package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// init is called when the program starts; no initial commands needed rn
func (m model) Init() tea.Cmd {
	return nil
}

// update handles all incoming events like keypresses/window resizes
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// this message tells us the size of the terminal window
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	// this message tells us a key was pressed
	case tea.KeyMsg:
		switch msg.String() {

		// these keys will exit the program
		case "ctrl+c", "esc":
			return m, tea.Quit

		// handle deleting the last character
		case "backspace":
			// only delete if there is actually something to delete
			if len(m.userInput) > 0 {
				m.userInput = m.userInput[:len(m.userInput)-1]
			}

		// handle all other regular typing keys
		default:
			// we only care about single character keypresses for now
			if len(msg.String()) == 1 {
				// if this is the very first keypress, start the timer
				if !m.started {
					m.started = true
					m.startTime = time.Now()
				}
				// add the typed character to our input tracker
				m.userInput += msg.String()
			}
		}
	}

	// return the updated model and no new commands
	return m, nil
}
