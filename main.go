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

	focusMinutes := usrinput.UsrInput()

	p := tea.NewProgram(pomodoro.FocusModel(focusMinutes)) //start from the initial model
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
