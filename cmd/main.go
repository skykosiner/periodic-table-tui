package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yonikosiner/perodic-table-tui/pkg/elements"
)

/*
I don't think I can do this with buble tea? I need to be able to have seperate
rows with elements, and these rows have different hights.

TODO: Look at lipgloss by charm to get the styles I want
*/

func main() {
	p := tea.NewProgram(elements.InitialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
