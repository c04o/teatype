package main

import (
	"fmt"
)

// view draws the user interface to the terminal screen
func (m model) View() string {
	// get the current words per minute
	wpm := m.calculateWPM()

	// create a simple formatted string to display the wpm
	wpmDisplay := fmt.Sprintf("wpm: %.0f", wpm)

	// build the main content string with standard formatting
	// use newlines to create vertical spacing
	content := fmt.Sprintf(
		"%s\n\n%s\n\ninput: %s\n\n(esc to quit)",
		wpmDisplay,
		m.targetText,
		m.userInput,
	)

	// return the plain text content
	// this will naturally render at the top left of the terminal
	return content
}
