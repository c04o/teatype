package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// main entry point of app
func main() {
	// create a new bubble tea program using initial model state
	// use the alt screen so it cleans up the terminal when we quit
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	// run the program and check if any fatal errors occurred
	if _, err := p.Run(); err != nil {
		// print error if failed
		fmt.Printf("error: %v", err)
		// exit with a non-zero status code to indicate failure
		os.Exit(1)
	}
}
