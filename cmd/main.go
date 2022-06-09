package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yonikosiner/perodic-table-tui/pkg/elements"
)

type item string

type model struct {
	list     list.Model
	items    []item
	choice   string
	quitting bool
}

func main() {
	elementss := elements.GetElementStr()

	var elmentArr []list.Item

	for i := 0; i < len(elementss); i++ {
		elmentArr = append(elmentArr, elements.Item(elementss[i]))
	}

	const defaultWidth = 20
	l := list.New(elmentArr, elements.ItemDelaget{}, defaultWidth, elements.ListHeight)

	l.Title = "Perdoic Table TUI"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = elements.TitleStyle
	l.Styles.PaginationStyle = elements.PaginationStyle
	l.Styles.HelpStyle = elements.HelpStyle

	m := elements.Model{List: l}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
