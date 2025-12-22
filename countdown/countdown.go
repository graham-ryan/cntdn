// Package countdown implements tea timers for a countdown timer
package countdown

import (
	"fmt"

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
	ti.Placeholder = "30" // idea: use previously used time
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
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case error:
		m.err = msg
		return m, nil
	}

	if m.textInput.Focused() {
		if msg == "" {
			m.textInput.Blur()
		}
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	} else {
		if msg == " " {
			m.timer.Init()
		}
		return m, nil
	}
}

func (m model) View() string {
	// The header
	header := "cntdn"

	// The footer
	footer := "\nesc • quit | enter • toggle | pace • toggle cntdn\n"

	var time string
	if m.timer.Running() {
		time = m.timer.View()
	}

	return fmt.Sprintf("%s\n\n%s\n%s\n\n%s\n", header, m.textInput.View(), time, footer)
}
