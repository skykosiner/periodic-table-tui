package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yonikosiner/perodic-table-tui/pkg/elements"
)

func main() {
	elementss := elements.GetElementStr()

	var elmentArr []list.Item

	for i := 0; i < len(elementss); i++ {
		elmentArr = append(elmentArr, elements.Item(elementss[i]))
	}

	const defaultWidth = 69
	l := list.New(elmentArr, elements.ItemDelaget{}, defaultWidth, elements.ListHeight)

	l.Title = "Perdoic Table TUI\nKey: Group 69 means it is in the middle of the\ntable where there is no group in the\nedexceul iGCSE periodic table"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = elements.TitleStyle
	l.Styles.PaginationStyle = elements.PaginationStyle
	l.Styles.HelpStyle = elements.HelpStyle

	m := elements.Model{List: l}

	if err := tea.NewProgram(m, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
