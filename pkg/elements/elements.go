package elements

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type element struct {
	Symbol            string   `json:"symbol"`
	Coordinates       []int    `json:"coordinates"`
	Names             []string `json:"names"`
	Valence           string   `json:"valence"`
	Neutrons          int      `json:"neutrons"`
	Protons           int      `json:"protons"`
	State             string   `json:"state"`
	Radioactivity     string   `json:"radioactivity"`
	Radius            string   `json:"radius"`
	Electronegativity string   `json:"electronegativity"`
	Density           string   `json:"density"`
	Melting           string   `json:"melting"`
	Boiling           string   `json:"boiling"`
	Discoverer        string   `json:"discoverer"`
	Year              string   `json:"year"`
	Specific_heat     string   `json:"specific_heat"`
	First_ionization  string   `json:"first_ionization"`
}

type model struct {
	elements []element
	cursor   int
	selected map[int]struct{}

	table []string
}

func InitialModel() model {
	jsonFile, err := os.Open("elements.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var el []element

	json.Unmarshal(byteValue, &el)

	return model{
		elements: el[0:3],

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.elements)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string
	for i, choice := range m.elements {
		var sideInfo element
		if m.cursor == i {
			sideInfo = choice
		}

		jsonStr, err := json.MarshalIndent(sideInfo, "", "  ")

		if err != nil {
			panic("There was an error changing the info to a json string")
		}

		if m.cursor != i {
			for _, element := range m.elements {
				if element.Symbol != sideInfo.Symbol {
					s += m.GenrateSymbol(element.Symbol)
				}
			}

		} else {
			s += fmt.Sprintf("%s\n%s\n", m.GenrateSymbol(sideInfo.Symbol), jsonStr)
		}
	}

	s += "\nPress q to quit.\n"
	return s
}

func (m *model) GenrateSymbol(symbol string) string {
	// groups := []int{1, 2, 3, 4, 5, 6, 7, 0}

	/*
	   Each element should look like this:
	   ╭─╮
	   │H│
	   ╰─╯
	*/

	var finalString string
	for y := 0; y < len(symbol); y++ {
		for x := 0; x < len(symbol); x++ {
			if y != 0 && x == 0 {
				finalString += "╭"
			}

			if y == 1 && x != 1 {
				finalString += "─"
			}

			if y == 1 && x != 1 {
				finalString += "─"
			}

			if y == 1 && x != 1 {
				finalString += "╮"
			}

			if y == len(symbol)-1 && x != len(symbol)-1 {
				finalString += "\n│"
			}

			if y == 1 && x != 1 {
				finalString += fmt.Sprintf("%s│", symbol)
			}

			if y != 0 && x == 0 {
				finalString += "\n╰"
			}

			if y == 1 && x != 1 {
				finalString += "─"
			}

			if y == 1 && x != 1 {
				finalString += "─"
			}

			if y == 1 && x != 1 {
				finalString += "╯\n"
			}
		}
	}

	return finalString
}
