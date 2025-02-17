package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylanbloodworth/pomodoro/internal/timer"
	"github.com/dylanbloodworth/pomodoro/internal/usrinput"
)

// main runs the pomodoro application
func main() {

	minutes, poms := usrinput.UsrInput() //Starts the user input screen and returns the form value.

	if minutes != 0 && poms != 0 {
		p := tea.NewProgram(timer.FocusModel(minutes, poms)) //starts the pomodoro timer
		if _, err := p.Run(); err != nil {
			fmt.Printf("There's been an error: %v", err)
			os.Exit(1)
		}

	}
}
