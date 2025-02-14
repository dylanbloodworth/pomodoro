package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type model struct {
	textInput textinput.Model
	usrInput  string
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		//then look at the key that was pressed
		switch msg.String() {
		//If the key was enter
		case "esc":
			//Return the model and quit
			return m, tea.Quit
		case "enter":
			m.usrInput = m.textInput.View()
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	form := fmt.Sprintf(
		"What am I typing?, \n\n%s",
		m.textInput.View(),
	)
	output := m.usrInput

	return form + "\n\n" + output
}

func InitialModel() model {
	ti := textinput.New()
	ti.Focus()

	return model{textInput: ti}
}

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
