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
	elements []element        // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func InitialModel() model {
	// Get each element from json file
	// Open our jsonFile
	jsonFile, err := os.Open("elements.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var el []element

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &el)

	return model{
		// Our shopping list is a grocery list
		elements: el[0:2],

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.elements)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	var s string

	// Iterate over our choices
	for i, choice := range m.elements {
		var sideInfo element
		if m.cursor == i {
			sideInfo = choice
		}

		// Render the row
		// EWWWWWWWW
		s += fmt.Sprintf("%s\n \t\t\t\t\tNames: %s\n\t\t\t\t\tValence: %s\n\t\t\t\t\tNuetrons: %d\n\t\t\t\t\tProtons: %d\n\t\t\t\t\tState: %s\n\t\t\t\t\tRadioactivity: %s\n\t\t\t\t\tRadius: %s\n\t\t\t\t\tElectronegativity: %s\n\t\t\t\t\tDensity: %s\n\t\t\t\t\tMelting: %s\n\t\t\t\t\tBoiling: %s\n\t\t\t\t\tDiscoverer: %s\n\t\t\t\t\tYear: %s\n\t\t\t\t\tSpecific_heat: %s\n\t\t\t\t\tFirst_ionization: %s\n", choice.Symbol, sideInfo.Names, sideInfo.Valence, sideInfo.Neutrons, sideInfo.Protons, sideInfo.State, sideInfo.Radioactivity, sideInfo.Radius, sideInfo.Electronegativity, sideInfo.Density, sideInfo.Melting, sideInfo.Boiling, sideInfo.Discoverer, sideInfo.Year, sideInfo.Specific_heat, sideInfo.First_ionization)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
