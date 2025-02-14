package main

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

// model contains a textinput which is a form to type.
type model struct {
	textInput textinput.Model
	err       error
}

func initialModel() model {
	ti := textinput.New()      // Creates the space to start typing
	ti.Placeholder = "Pikachu" // Creates a placeholder text in the form
	ti.Focus()                 // Focuses the model for it to receive key inputs
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,  // The initial state of the model has the intial text input
		err:       nil, // Err set to nil? Don't know why?
	}
}

// Seems like this just makes the cursor blink at the start without having to press any buttons?
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg) //updates the key presses
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		m.textInput.View(), //View port for the text input
		"(esc to quit)",
	) + "\n"
}
