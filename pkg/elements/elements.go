package elements

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type element struct {
	Symbol            string   `json:"symbol"`
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

const ListHeight = 20

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

func (i Item) FilterValue() string {
	return string(i)
}

type ItemDelaget struct{}

func (d ItemDelaget) Height() int                               { return 1 }
func (d ItemDelaget) Spacing() int                              { return 0 }
func (d ItemDelaget) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d ItemDelaget) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := ItemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return SelectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

type Item string

type Model struct {
	List     list.Model
	Elements []Item
	Choice   string
	Quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.List.FilterState() == list.Filtering {
			break
		}

		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit
		case "enter":
			i, ok := m.List.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.Choice != "" {
		element := GetElementbySymbol(m.Choice)

		byte, err := json.MarshalIndent(element, " ", " ")

		if err != nil {
			log.Fatal("There was an error settitg the element info to a json string", err)
		}

		return QuitTextStyle.Render(string(byte))
	}

	if m.Quitting {
		return QuitTextStyle.Render("You don't want to look at the elements in the perdoic table?")
	}

	return "\n" + m.List.View()
}

func GetElementStr() []string {
	jsonFile, err := os.Open("elements.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var el []element
	json.Unmarshal(byteValue, &el)

	var symbols []string

	for i := 0; i < len(el); i++ {
		symbol := el[i].Symbol
		symbols = append(symbols, symbol)
	}

	return symbols
}

func GetElements() []element {
	jsonFile, err := os.Open("elements.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var el []element
	json.Unmarshal(byteValue, &el)

	return el
}

func GetElementbySymbol(symbol string) element {
	elements := GetElements()
	var element element

	for _, el := range elements {
		if el.Symbol == symbol {
			element = el
		}
	}

	return element
}
