package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"time"
)

// Create a Msg to update the module
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
		case "enter":
			//Return the model and quit
			return m, tea.Quit
		}

	// Update every second
	case TickMsg:

		//Quit Program after 10 seconds
		if m.curTime == 0 {
			return m, tea.Quit
		} else {
			m.curTime -= time.Second // update timer every second

			//Check if 15 seconds has passed to update the progress bar
			if m.curTime%(15*time.Second) == 0 {
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
	s += fmt.Sprintf("Time: %v  ", m.curTime)
	s += m.progress
	return s
}

func InitialModel(curTime time.Duration, totalPoms int8) model {
	return model{curTime: curTime, poms: 0, totalPoms: totalPoms, progress: ""}
}

func main() {
	p := tea.NewProgram(InitialModel(time.Minute, 10))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}
