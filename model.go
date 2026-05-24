package main

import "time"

// hold the entire state of the app
type model struct {
	targetText string
	userInput  string
	startTime  time.Time
	started    bool
	width      int
	height     int
}

// create the starting state of app
func initialModel() model {
	return model{
		targetText: "the quick brown fox jumps over the lazy dog",
		userInput:  "",
	}
}

// calculates the current wpm based on input length
func (m model) calculateWPM() float64 {
	// if input hasn't started, the wpm is 0
	if !m.started {
		return 0
	}

	// figure out how many minutes have passed since the first keystroke
	elapsed := time.Since(m.startTime).Minutes()

	// prevent division by zero in the very first ms duh
	if elapsed <= 0 {
		return 0
	}

	// a standard typing "word" is exactly 5 characters long
	return (float64(len(m.userInput)) / 5.0) / elapsed
}
