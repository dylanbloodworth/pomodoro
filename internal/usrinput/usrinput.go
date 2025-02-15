package usrinput

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var usrInput string

type model struct {
	textInput textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
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
			usrInput = strings.TrimSpace(strings.TrimPrefix(m.textInput.View(), "> "))
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	formHeading := "How many minutes will your focus sessions be?\n"
	form := fmt.Sprint(m.textInput.View())
	usrInput = m.textInput.Value()

	return formHeading + form
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Please enter an integer value (e.g. 25)"
	ti.Focus()
	ti.Cursor.SetMode(cursor.CursorHide)

	return model{textInput: ti}
}

func UsrInput() int {
	p := tea.NewProgram(InitialModel())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}

	i, err := strconv.Atoi(usrInput)
	if err != nil {
		panic(err)
	}

	return i
}
