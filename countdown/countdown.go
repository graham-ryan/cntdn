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
	ti.CharLimit = 20
	ti.Width = 20

	return model{
		timer:     timer.New(5 * time.Minute),
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
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.TimeoutMsg:
		notifyErr := notify()
		if notifyErr != nil {
			m.err = notifyErr
		}
		cmd = m.textInput.Focus()
		return m, cmd

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyEnter:
			if m.textInput.Focused() {
				var initialTimeout string
				if m.textInput.Value() != "" {
					initialTimeout = m.textInput.Value()
				} else {
					initialTimeout = m.textInput.Placeholder
					m.textInput.SetValue(m.textInput.Placeholder)
				}

				parsedTimeout, err := parseTime(initialTimeout)
				if err != nil {
					m.err = err
					return m, nil
				}

				m.textInput.Blur()
				m.timer = timer.NewWithInterval(parsedTimeout, 20*time.Millisecond)
				cmd = m.timer.Init()

				return m, cmd
			} else {
				cmd = m.timer.Toggle()
				return m, cmd
			}

		case tea.KeyRunes, tea.KeyBackspace:
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

	var timeStr string
	if m.timer.Running() {
		timeStr = "\n" + m.timer.Timeout.Round(time.Second).String()
	} else if !m.timer.Timedout() {
		timeStr = "\n" + m.timer.Timeout.Round(time.Second).String() + "\t(STOPPED)"
	} else {
		timeStr = "\n"
	}

	return fmt.Sprintf("%s\n\n%s\n%s\n\n%s\n", header, m.textInput.View(), timeStr, footer)
}
