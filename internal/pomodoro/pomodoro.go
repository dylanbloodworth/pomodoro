// The pomodoro package creates an application for implementing the pomodoro
// technique (https://en.wikipedia.org/wiki/Pomodoro_Technique). It renders a
// TUI that runs a pomodoro timer with user defined timing conditions.
package pomodoro

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// The statusCodes map translates the int8 pomodoro.status field into a string for
// rendering on the TUI.
var statusCodes map[int8]string = map[int8]string{
	0: "Focusing",
	1: "Short Break",
	2: "Long Break",
}

// The pomodoro structure holds information about the state of the pomodoro application.
type pomodoro struct {
	curTime   time.Duration // The curTime field acts as the timer for the current status.
	poms      int8          // The poms field indicates the number of poms cycle currently completed.
	totalPoms int8          // The totalPoms field indicates the total number of pomodoro cycles to complete.
	status    int8          // The status field tracks the current status of the pomodoro cycle.
	progress  string        // The progress field is a string representation of the progress bar to be rendered on the TUI.
}

// FocusModel initializes a pomodoro model with some initial state. It returns
// the pomodoro model starting in the focused status.
//
// The eventual goal should be to have curTime, poms, and totalPoms be user defined
// fields before the pomodoro timer starts.
func FocusModel() pomodoro {
	return pomodoro{
		curTime:   15 * time.Second,
		poms:      0,
		totalPoms: 4,
		status:    0,
		progress:  ""}
}

// Create a Msg to update the module based on time
type TickMsg time.Time

// TickEvery() returns a TickMsg every second. This updates the TUI every second.
func TickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Init is the first function to be called. It returns the TickEvery function to
// tick every second.
func (p pomodoro) Init() tea.Cmd {
	return TickEvery()
}

// Update returns an updated state of the model based on internal logic of the program.
//
// The pomodoro timer checks if the escape button has been pressed to quit the program.
// It handles switching between statuses after the timer ends and tracks the completed
// pomodoro cycles. The logic is based off a simple pomodoro timer structure. Complete
// 4 focus cycles with a short break in between then complete 1 long break.
func (p pomodoro) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Is the msg sent to update a key press?
	case tea.KeyMsg:
		//then look at the key that was pressed
		switch msg.String() {

		//If the key was enter
		case "esc":
			//Return the model and quit
			return p, tea.Quit
		}

	// Update every second
	case TickMsg:

		switch p.curTime {
		case 0:
			p.progress = " "

			switch p.status {
			case 0:
				p.poms += 1
				if p.poms == p.totalPoms {
					return p, tea.Quit
				}

				if p.poms%4 == 0 && p.poms != 0 {
					p.status = 2
					p.curTime = 10 * time.Second
				} else {
					p.status = 1
					p.curTime = 5 * time.Second
				}

			default:
				p.status = 0
				p.curTime = 15 * time.Second
			}
			return p, TickEvery()

		default:
			p.curTime -= time.Second
			if p.curTime%(time.Second) == 0 {
				p.progress += "%"
			}
			return p, TickEvery()
		}
	}
	return p, nil
}

// View returns a string which represents the rendering of the TUI.
//
// It renders a header with the currently complete pomodoro cycles,
// the total amount of cycles to be completed, the status of the
// cycle, a countdown timer for the current pomodoro cycle, and
// the status of the pomodoro cycle, and the progress bar. All of
// pomodoro type fields are rendered.
func (m pomodoro) View() string {
	s := "\n ------- Pomodoro Timer -------- \n"
	s += fmt.Sprintf(" ---- Poms Complete : %d / %d ---- \n", m.poms, m.totalPoms)
	s += fmt.Sprintf(" (%v) | Time: %v  ", statusCodes[m.status], m.curTime)
	s += m.progress
	s += "\n"
	return s
}
