package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dylanbloodworth/charm/internal/pomodoro"
	"os"
)

// main runs the pomodoro application
func main() {
	p := tea.NewProgram(pomodoro.FocusModel()) //start from the initial model
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
