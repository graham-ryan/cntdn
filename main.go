package main

import (
	"fmt"
	"os"

	"graham-ryan/cntdn/countdown"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(countdown.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
