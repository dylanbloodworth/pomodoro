package main

import (
	"fmt"
	// "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

var poms int8
var totalPoms int8

var statusCodes map[int8]string = map[int8]string{
	0: "Focusing",
	1: "On Short Break",
	2: "On Long Break",
}

// Run the application
func main() {
	p := tea.NewProgram(FocusModel(15 * time.Second)) //start from the initial model
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func FocusModel(curTime time.Duration) model {
	return model{curTime: curTime, poms: poms, totalPoms: totalPoms, status: 0, progress: ""}
}

// Create a Msg to update the module based on time
type TickMsg time.Time

func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

type model struct {
	curTime   time.Duration // current time on the timer
	poms      int8          //number of pomodoros completed
	totalPoms int8          //total amount of poms that want to be run
	status    int8          //Status of the model (compare to status codes)
	progress  string        //String representing the progress bar
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is the msg sent to update a key press?
	case tea.KeyMsg:
		//then look at the key that was pressed
		switch msg.String() {

		//If the key was enter
		case "esc":
			//Return the model and quit
			return m, tea.Quit
		}

	// Update every second
	case TickMsg:

		switch m.curTime {
		case 0:
			m.progress = " "

			switch m.status {
			case 0:
				m.poms += 1
				if m.poms%4 == 0 && m.poms != 0 {
					m.status = 2
					m.curTime = 10 * time.Second
				} else {
					m.status = 1
					m.curTime = 5 * time.Second
				}

			case 1:
				m.status = 0
				m.curTime = 15 * time.Second

			case 2:
				m.status = 0
				m.curTime = 15 * time.Second
			}
			return m, tickEvery()

		default:
			m.curTime -= time.Second
			if m.curTime%(time.Second) == 0 {
				m.progress += "%"
			}
			return m, tickEvery()
		}
	}
	return m, nil
}

func (m model) View() string {

	s := "\n ------- Pomodoro Timer -------- \n"
	s += fmt.Sprintf(" ---- Poms Complete : %d / %d ---- \n", m.poms, m.totalPoms)
	s += fmt.Sprintf(" | Time: %v  ", m.curTime)
	s += m.progress
	s += "\n"
	return s
}
