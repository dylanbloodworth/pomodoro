package usrinput

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// usrInput is the global variable to log what the value of the textinput form.
var usrInput string

// model type for holding a textInput model
type model struct {
	textInput   textinput.Model
	clearScreen bool
}

// Init returns nil so no command is run at the beginning of the model
func (m model) Init() tea.Cmd {
	return nil
}

// Update returns an updated state of the textInput model. It tracks
// the text input into the field and assigns the value of the global
// usrInput variable.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		//then look at the key that was pressed
		switch msg.String() {
		//If the key was escape
		case "esc":
			//Return quit the form input without doing anything else.
			return model{clearScreen: true}, tea.Quit
		case "enter":
			//track the final usrInput as a string and quit the program
			usrInput = m.textInput.Value()
			return model{clearScreen: true}, tea.Quit
		}
	}

	//handles user keystrokes in the form
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the form by returning a string. Most of the rendering occurs in the textinput
// bubbles UI component
func (m model) View() string {

	if m.clearScreen {
		return ""
	}

	formHeading := "How many minutes will your focus sessions be?\n"
	form := fmt.Sprint(m.textInput.View())

	return lipgloss.NewStyle().Padding(1, 0, 0, 2).Foreground(lipgloss.Color("231")).Render(formHeading + form)
}

// InitialModel defines the configs of the textinput model and returns
// the type model defined at the start of the program.
func InitialModel() model {
	ti := textinput.New()                                                  // declare new textinput. See textinput.New() definition to set configs
	ti.Placeholder = "Please enter an integer value (e.g. 25)"             // suggests a recommended pomodoro time
	ti.Focus()                                                             // allows the form to receive keystrokes
	ti.Cursor.SetMode(cursor.CursorHide)                                   // hides the cursor in the form
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("177")) // change the color of the '>' in the prompt

	return model{textInput: ti, clearScreen: false}
}

// UsrInput runs the program that asks the user to input their desired pomodoro focus time in minutes.
// It returns the desired time in minutes as an integer.
func UsrInput() int {
	p := tea.NewProgram(InitialModel())

	_, err := p.Run()
	if err != nil {
		panic(err)
	}

	if usrInput != "" {
		minutes, err := strconv.Atoi(usrInput)
		if err != nil {
			panic(err)
		}
		return minutes
	} else {
		fmt.Print("No Input Provided")
		return 0
	}
}
