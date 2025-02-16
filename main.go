package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylanbloodworth/charm/internal/pomodoro"
	"github.com/dylanbloodworth/charm/internal/usrinput"
)

// main runs the pomodoro application
func main() {

	focusMinutes := usrinput.UsrInput() //Starts the user input screen and returns the form value.

	if focusMinutes != 0 {
		p := tea.NewProgram(pomodoro.FocusModel(focusMinutes)) //starts the pomodoro timer
		if _, err := p.Run(); err != nil {
			fmt.Printf("There's been an error: %v", err)
			os.Exit(1)
		}

	}
}
