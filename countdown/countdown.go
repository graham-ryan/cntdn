// Package countdown implements tea timers for a countdown timer
package countdown

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/timer"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	timer     timer.Model
	textInput textinput.Model
	err       error
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "5m" // idea: use previously used time
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		cmd = m.textInput.Focus()
		return m, cmd

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.textInput.Focused() {
				if m.textInput.Value() != "" {
					m.textInput.Blur()
					m.timer = timer.NewWithInterval(5*time.Minute, 1*time.Second)
					cmd = m.timer.Init()
				}
				return m, cmd
			} else {
				cmd = m.timer.Toggle()
				return m, cmd
			}
		case tea.KeyRunes:
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}

	case error:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// The header
	header := "cntdn"

	// The footer
	footer := "\nesc • quit | enter • toggle cntdn\n"

	var time string
	if m.timer.Running() {
		time = "\n" + m.timer.View()
	} else if !m.timer.Timedout() {
		time = "\n" + m.timer.View() + "\t(STOPPED)"
	} else {
		time = "\n"
	}

	return fmt.Sprintf("%s\n\n%s\n%s\n\n%s\n", header, m.textInput.View(), time, footer)
}
